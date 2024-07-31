package capturer

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func RequestScreenshot(addr, outDirPath string) {
	resp, err := http.Get(addr)
	if err != nil {
		log.Printf("Error requesting screenshot from %s: %v", addr, err)
		return
	}
	defer resp.Body.Close()
	// Check for successful response status code
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Server returned status code %d while fetching screenshot from %s", resp.StatusCode, addr)
		return
	}

	// Generate a unique filename with the current timestamp
	filename := fmt.Sprintf(
		"screenshot_%d-%02d-%02dT%02d-%02d-%02d.png",
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

	fmt.Println("Image saved successfully to", fullPath)
}
