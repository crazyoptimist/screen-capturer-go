package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"time"

	"github.com/pion/mdns/v2"

	"screencapturer/internal/config"
	"screencapturer/internal/constant"
	"screencapturer/internal/domain/model"
	"screencapturer/internal/infrastructure/capturer"
	"screencapturer/internal/infrastructure/mdnsserver"
	"screencapturer/internal/infrastructure/server"
)

const REQUEST_INTERVAL_IN_SECONDS = 5

var outDirPath string
var wg sync.WaitGroup

func main() {
	flag.StringVar(&outDirPath, "dir", "", "Output directory name")
	flag.Parse()

	if outDirPath == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln("Couldn't find the user home dir.")
		}
		outDirPath = filepath.Join(userHomeDir, "Downloads")
	}

	// Initialize the database
	if _, err := config.InitDB(); err != nil {
		log.Fatalln("Database initialization failed: ", err)
	}

	// Request screenshots from all servers periodically
	go func() {
		for range time.Tick(REQUEST_INTERVAL_IN_SECONDS * time.Second) {
			// Get all the registered computers
			var computers = []model.Computer{}
			if err := config.DB.Limit(256).Find(&computers).Error; err != nil {
				log.Fatalln("Database querying failed: ", err)
			}

			for _, pc := range computers {
				addr := pc.GetEndpoint()
				if addr == "" || !pc.IsActive {
					continue
				}

				outPath := filepath.Join(outDirPath, pc.Name)

				wg.Add(1)
				go func(addr, outPath string) {
					defer wg.Done()
					capturer.RequestScreenshot(addr, outPath)
				}(addr, outPath)
			}
			wg.Wait()
		}
	}()

	// Periodically scan computers in the LAN
	mDNSServer, err := mdnsserver.CreateMDNSServer(true, false, &mdns.Config{})
	if err != nil {
		log.Fatalln("Creating a mDNS instance failed: ", err)
	}
	go func(*mdns.Conn) {
		for range time.Tick(constant.SCAN_NETWORK_INTERVAL_IN_SECONDS * time.Second) {
			capturer.Scan(mDNSServer)
		}
	}(mDNSServer)

	// Expose Web/UI
	httpServer := server.NewServer()
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalln("Web server startup failed: ", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalln("HTTP server shutdown failed: ", err)
	}
	log.Println("Graceful shutdown finished.")
}
