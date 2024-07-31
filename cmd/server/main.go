package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"screencapturer/internal/constant"
	"screencapturer/internal/infrastructure/capturer"
	"screencapturer/internal/infrastructure/mdnsserver"
)

const CAPTURE_INTERVAL_IN_SECONDS = 5

var vhost string

func main() {
	flag.StringVar(&vhost, "vhost", "", "Virtual host name")
	flag.Parse()

	if vhost == "" {
		fmt.Println("Virtual host name must be set.")
		return
	}

	// Generate a merged screenshot periodically
	go func() {
		for range time.Tick(CAPTURE_INTERVAL_IN_SECONDS * time.Second) {
			capturer.CaptureMergedScreen()
		}
	}()

	// Serve the screenshot over HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")

		imgPath := capturer.GetMergedScreenFilePath()
		data, err := os.ReadFile(imgPath)
		if err != nil {
			fmt.Fprintf(w, "Error reading screen file: %v", err)
			return
		}

		_, err = w.Write(data)
		if err != nil {
			fmt.Fprintf(w, "Error writing screen file: %v", err)
		}
	})

	// Publish via mdns server
	if err := mdnsserver.ListenMDNS(vhost); err != nil {
		log.Fatalln("Listening on MDNS queries failed: ", err)
	}

	// Run the web server
	webServerHost := fmt.Sprintf(":%d", constant.SERVER_WEB_PORT)
	fmt.Printf("Server starting on %s\n", webServerHost)
	if err := http.ListenAndServe(webServerHost, nil); err != nil {
		log.Fatalln("Listening on web requests failed: ", err)
	}
}
