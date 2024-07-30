package mdnsserver

import (
	"net"

	"github.com/pion/mdns/v2"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// This function creates a MDNS server that listens to incoming MDNS queries.
// It's not blocking.
//
// # Params
//
// vhost - A virtual host name that can identify this server
func ListenMDNS(vhost string) error {
	addr4, err := net.ResolveUDPAddr("udp4", mdns.DefaultAddressIPv4)
	if err != nil {
		return err
	}

	addr6, err := net.ResolveUDPAddr("udp6", mdns.DefaultAddressIPv6)
	if err != nil {
		return err
	}

	l4, err := net.ListenUDP("udp4", addr4)
	if err != nil {
		return err
	}

	l6, err := net.ListenUDP("udp6", addr6)
	if err != nil {
		return err
	}

	_, err = mdns.Server(ipv4.NewPacketConn(l4), ipv6.NewPacketConn(l6), &mdns.Config{
		LocalNames: []string{vhost},
	})
	if err != nil {
		return err
	}

	return nil
}
