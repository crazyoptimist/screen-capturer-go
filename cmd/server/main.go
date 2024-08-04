package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"screencapturer/internal/constant"
	"screencapturer/internal/infrastructure/capturer"
	"screencapturer/internal/infrastructure/mdnsserver"
)

const CAPTURE_INTERVAL_IN_SECONDS = 11

var vhost string

func main() {
	flag.StringVar(&vhost, "vhost", "", "Virtual host name")
	flag.Parse()

	if vhost == "" {
		fmt.Println("Virtual host name must be set.")
		return
	}

	// Capture screen if requested,
	// serve the screenshot over HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")

		if err := capturer.CaptureMergedScreen(); err != nil {
			log.Printf("Error capturing the merged screen: %v", err)
			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintf(w, "Error capturing the merged screen: %v", err)
			return
		}

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
		return
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
