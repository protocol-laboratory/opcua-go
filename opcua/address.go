package opcua

import (
	"net"
	"strconv"
)

type Address struct {
	// Host domain name or ipv4, ipv6 address
	Host string
	// Port service port
	Port int
}

func (addr Address) Addr() string {
	return net.JoinHostPort(addr.Host, strconv.Itoa(addr.Port))
}
