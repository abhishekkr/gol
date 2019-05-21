package golnet

import (
	"net"
)

func IsPortOpen(portNumber string) bool {
	ln, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		return true
	}
	ln.Close()
	return false
}

func IsPortClosed(portNumber string) bool {
	return !IsPortOpen(portNumber)
}
