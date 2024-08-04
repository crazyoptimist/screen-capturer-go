package capturer

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"screencapturer/internal/config"
	"screencapturer/internal/constant"
	"screencapturer/internal/domain/model"
	"screencapturer/pkg/utils"
	"time"

	"github.com/pion/mdns/v2"
)

const HTTP_REQUEST_TIMEOUT_IN_SECONDS = 3

func RequestScreenshot(addr, outDirPath string) {
	if err := utils.CreateDirIfNotExists(outDirPath); err != nil {
		log.Println("Error creating the output directory failed: ", err)
		return
	}

	client := http.Client{
		Timeout: HTTP_REQUEST_TIMEOUT_IN_SECONDS * time.Second, // Set timeout to 5 seconds
	}

	resp, err := client.Get(addr)
	if err != nil {
		log.Printf("Error requesting screenshot from %s: %v", addr, err)
		return
	}
	defer resp.Body.Close()

	// Check for successful response status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		log.Printf("Status code %d from %s: %s", resp.StatusCode, addr, string(bodyBytes))
		return
	}

	// Generate a unique filename with the current timestamp
	filename := fmt.Sprintf(
		"screenshot_%d-%02d-%02dT%02d-%02d-%02d.jpeg",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second(),
	)

	// Construct the full path where the screenshot will be saved
	fullPath := filepath.Join(outDirPath, filename)

	// Create a new file for saving the image
	outFile, err := os.Create(fullPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outFile.Close()

	// Copy the image data from the response to the local file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		fmt.Println("Error saving image:", err)
		return
	}

	// log.Println("Image saved successfully to", fullPath)
}

// This function queries all the registered computers
// and updates the active status with the found ip addresses
func Scan(mDNSServer *mdns.Conn) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		constant.SCAN_NETWORK_INTERVAL_IN_SECONDS*time.Second,
	)
	defer cancel()

	var computers = []model.Computer{}
	if err := config.DB.Limit(256).Find(&computers).Error; err != nil {
		log.Fatalln("Database querying failed: ", err)
		return err
	}

	for _, pc := range computers {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			_ = findComputer(mDNSServer, &pc)
		}
	}
	return nil
}

func findComputer(mDNSServer *mdns.Conn, pc *model.Computer) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		constant.FIND_PC_TIMEOUT_IN_SECONDS*time.Second,
	)
	defer cancel()

	_, ipAddr, err := mDNSServer.QueryAddr(ctx, pc.Name)
	if err != nil {
		// Use map here because GORM will only update non-zero fields if struct is used
		_ = config.DB.Model(pc).Updates(map[string]interface{}{"is_active": false}).Error
		return fmt.Errorf("mDNS: querying [%s] failed: %v\n", pc.Name, err)
	}

	if err := config.DB.Model(pc).Updates(model.Computer{
		IsActive:  true,
		IPAddress: fmt.Sprintf("%s", ipAddr),
	}).Error; err != nil {
		return err
	}

	return nil
}
