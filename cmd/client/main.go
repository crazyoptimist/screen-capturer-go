package main

import (
	"context"
	"fmt"

	"github.com/pion/mdns/v2"

	"screencapturer/internal/mdnsserver"
)

func main() {
	server, err := mdnsserver.CreateMDNSServer(true, false, &mdns.Config{})
	if err != nil {
		panic(err)
	}

	_, ipAddr, err := server.QueryAddr(context.TODO(), "snail")
	if err != nil {
		fmt.Println("Querying a vhost failed: ", err)
	}
	fmt.Println(ipAddr)
}
