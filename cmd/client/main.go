package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/pion/mdns/v2"

	"screencapturer/internal/constant"
	"screencapturer/internal/mdnsserver"
)

const REQUEST_INTERVAL_IN_SECONDS = 6

var wg sync.WaitGroup

func main() {
	var serverAddresses = []string{}

	server, err := mdnsserver.CreateMDNSServer(true, false, &mdns.Config{})
	if err != nil {
		panic(err)
	}

	_, ipAddr, err := server.QueryAddr(context.TODO(), "snail")
	if err != nil {
		fmt.Println("Querying a vhost failed: ", err)
	}
	httpAddr := fmt.Sprintf("http://%s:%d", ipAddr, constant.WEB_SERVER_PORT)
	serverAddresses = append(serverAddresses, httpAddr)

	// Request screenshots from all servers periodically
	for range time.Tick(REQUEST_INTERVAL_IN_SECONDS * time.Second) {
		wg.Add(len(serverAddresses))
		for _, addr := range serverAddresses {
			go func(addr string) {
				defer wg.Done()
				requestScreenshot(addr)
			}(addr)
		}
		wg.Wait()
	}
}

func requestScreenshot(addr string) {
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
	fullPath := filepath.Join("/tmp", filename)

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
