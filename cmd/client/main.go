package main

import (
	"context"
	"fmt"
	"net"

	"github.com/pion/mdns/v2"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

func main() {
	var useV4 = true
	var useV6 = false

	var packetConnV4 *ipv4.PacketConn
	if useV4 {
		addr4, err := net.ResolveUDPAddr("udp4", mdns.DefaultAddressIPv4)
		if err != nil {
			panic(err)
		}

		l4, err := net.ListenUDP("udp4", addr4)
		if err != nil {
			panic(err)
		}

		packetConnV4 = ipv4.NewPacketConn(l4)
	}

	var packetConnV6 *ipv6.PacketConn
	if useV6 {
		addr6, err := net.ResolveUDPAddr("udp6", mdns.DefaultAddressIPv6)
		if err != nil {
			panic(err)
		}

		l6, err := net.ListenUDP("udp6", addr6)
		if err != nil {
			panic(err)
		}

		packetConnV6 = ipv6.NewPacketConn(l6)
	}

	server, err := mdns.Server(packetConnV4, packetConnV6, &mdns.Config{})
	if err != nil {
		panic(err)
	}
	answer, src, err := server.QueryAddr(context.TODO(), "snail")
	if err != nil {
		fmt.Println("Querying a vhost failed: ", err)
	}
	fmt.Println(answer)
	fmt.Println(src)
}
