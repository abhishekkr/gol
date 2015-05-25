package golnet

import (
	"net"
)

func IsPortOpen(port_number string) bool {
	ln, err := net.Listen("tcp", ":"+port_number)
	if err != nil {
		return true
	}
	ln.Close()
	return false
}

func IsPortClosed(port_number string) bool {
	return !IsPortOpen(port_number)
}
