package capturer

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/pion/mdns/v2"

	"screencapturer/internal/config"
	"screencapturer/internal/constant"
	"screencapturer/internal/domain/model"
	"screencapturer/pkg/utils"
)

const CAPTURE_REQUEST_TIMEOUT_IN_SECONDS = 2
const MAX_RETRY_REQUEST_COUNT = 3
const RETRY_COOLDOWN_IN_MILLISECONDS = 500

func sendRequest(addr string) (*http.Response, error) {
	client := http.Client{
		Timeout: CAPTURE_REQUEST_TIMEOUT_IN_SECONDS * time.Second,
	}

	resp, err := client.Get(addr)
	if err != nil {
		return nil, fmt.Errorf("Error requesting screenshot from %s: %v", addr, err)
	}

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("Error code %d from %s: %s", resp.StatusCode, addr, string(bodyBytes))
	}

	return resp, nil
}

func saveResponseToImage(resp *http.Response, outDirPath string) error {
	if err := utils.CreateDirIfNotExists(outDirPath); err != nil {
		return fmt.Errorf("Error creating the output directory failed: %v", err)
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
		return fmt.Errorf("Error creating file: %v", err)
	}
	defer outFile.Close()

	// Copy the image data from the response to the local file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("Error saving image: %v", err)
	}

	return nil
}

func RequestScreenshot(addr, outDirPath string) error {
	var retry int8

	for retry = 0; retry < MAX_RETRY_REQUEST_COUNT; retry++ {
		resp, err := sendRequest(addr)
		if err != nil {
			time.Sleep(RETRY_COOLDOWN_IN_MILLISECONDS * time.Millisecond)
			continue
		} else {
			defer resp.Body.Close()
			if err = saveResponseToImage(resp, outDirPath); err != nil {
				return err
			}
			break
		}
	}

	if retry == MAX_RETRY_REQUEST_COUNT {
		return errors.New("Retry request exhausted.")
	}

	return nil
}

// This function queries all the registered computers
// and updates the IP addresses with the results
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
			ipAddress, err := findIPAddress(mDNSServer, &pc)
			if err != nil {
				// Use map here because GORM will only update non-zero fields if struct is used
				_ = config.DB.Model(pc).Updates(map[string]interface{}{"is_active": false}).Error
				log.Printf("Finding IP address failed for %s: %v", pc.Name, err)
				continue
			}

			if err := config.DB.Model(pc).Updates(model.Computer{
				IPAddress: ipAddress,
			}).Error; err != nil {
				log.Printf("Updating IP address failed: %v", err)
				continue
			}
		}
	}
	return nil
}

func findIPAddress(mDNSServer *mdns.Conn, pc *model.Computer) (string, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		constant.FIND_PC_TIMEOUT_IN_SECONDS*time.Second,
	)
	defer cancel()

	_, ipAddr, err := mDNSServer.QueryAddr(ctx, pc.Name)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", ipAddr), nil
}
