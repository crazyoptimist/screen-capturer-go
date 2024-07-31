package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"time"

	"github.com/pion/mdns/v2"

	"screencapturer/internal/config"
	"screencapturer/internal/domain/model"
	"screencapturer/internal/infrastructure/capturer"
	"screencapturer/internal/infrastructure/mdnsserver"
	"screencapturer/internal/infrastructure/server"
)

type CaptureServer struct {
	Name     string
	Endpoint string
}

const REQUEST_INTERVAL_IN_SECONDS = 60

var outDirPath string
var help bool
var wg sync.WaitGroup

func main() {
	flag.StringVar(&outDirPath, "dir", "", "Output directory name")
	flag.BoolVar(&help, "help", false, "Show help")
	flag.Parse()

	if help {
		fmt.Println(`
--dir <Output directory name> Stores the downloaded images to the output dir.
      It defaults to the Downloads folder if not provided.
--help Shows the usage guide`)
		return
	}

	if outDirPath == "" {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln("Couldn't get the user home dir.")
		}
		outDirPath = filepath.Join(userHomeDir, "Downloads")
	}

	// Create a mDNS server instance for querying
	mDNSServer, err := mdnsserver.CreateMDNSServer(true, false, &mdns.Config{})
	if err != nil {
		panic(err)
	}

	// Initialize the database
	if _, err := config.InitDB(); err != nil {
		panic(err)
	}

	var captureServers = []CaptureServer{}
	var computers = []model.Computer{}
	if err := config.DB.Limit(256).Find(&computers).Error; err != nil {
		log.Fatalln("Database querying failed: ", err)
	}

	for _, computer := range computers {
		ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
		res, ipAddr, err := mDNSServer.QueryAddr(ctx, computer.Name)
		if err != nil {
			fmt.Printf("mDNS: querying [%s] failed: %v\n", computer.Name, err)
		}
		fmt.Println(res)

		endpoint := fmt.Sprintf("http://%s:%d", ipAddr, config.WEB_SERVER_PORT)
		captureServers = append(captureServers, CaptureServer{Name: computer.Name, Endpoint: endpoint})
	}
	fmt.Println(captureServers)

	// Request screenshots from all servers periodically
	go func() {
		for range time.Tick(REQUEST_INTERVAL_IN_SECONDS * time.Second) {
			wg.Add(len(captureServers))
			for _, captureServer := range captureServers {
				addr := captureServer.Endpoint
				outPath := filepath.Join(outDirPath, captureServer.Name)

				go func(addr string) {
					defer wg.Done()
					capturer.RequestScreenshot(addr, outPath)
				}(addr)
			}
			wg.Wait()
		}
	}()

	// Expose Web/UI
	httpServer := server.NewServer()
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Web server startup failed: %v", err)
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
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}
	log.Println("Graceful shutdown finished.")
}
