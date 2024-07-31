package mdnsserver

import (
	"context"
	"fmt"
	"net"
	"net/netip"

	"github.com/pion/mdns/v2"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

func CreateMDNSServer(useV4, useV6 bool, config *mdns.Config) (*mdns.Conn, error) {
	var packetConnV4 *ipv4.PacketConn
	if useV4 {
		addr4, err := net.ResolveUDPAddr("udp4", mdns.DefaultAddressIPv4)
		if err != nil {
			return nil, err
		}

		l4, err := net.ListenUDP("udp4", addr4)
		if err != nil {
			return nil, err
		}

		packetConnV4 = ipv4.NewPacketConn(l4)
	}

	var packetConnV6 *ipv6.PacketConn
	if useV6 {
		addr6, err := net.ResolveUDPAddr("udp6", mdns.DefaultAddressIPv6)
		if err != nil {
			return nil, err
		}

		l6, err := net.ListenUDP("udp6", addr6)
		if err != nil {
			return nil, err
		}

		packetConnV6 = ipv6.NewPacketConn(l6)
	}

	return mdns.Server(packetConnV4, packetConnV6, config)
}

// vhost - A virtual host name that can identify this server
func ListenMDNS(vhost string) error {
	_, err := CreateMDNSServer(true, true, &mdns.Config{
		LocalNames: []string{vhost},
	})
	if err != nil {
		return err
	}
	return nil
}

func FindServer(server *mdns.Conn, vhost string) (*netip.Addr, error) {
	_, ipAddr, err := server.QueryAddr(context.Background(), vhost)
	if err != nil {
		return nil, fmt.Errorf("Querying for [%s] failed: %v", vhost, err)
	}

	return &ipAddr, nil
}
