package golzmq

import (
	"fmt"
	"time"
)

/* Structure to sustain ZMQ Proxy configuration */
type ZmqProxyConfig struct {
	DestinationIP    string
	DestinationPorts []int
	SourceIP         string
	SourcePorts      []int
}

/* Create a Proxy connection for given ZmqProxyConfig */
func ZmqProxySocket(proxy ZmqProxyConfig) error {
	chan_source := make(chan []byte, 5)
	chan_destination := make(chan []byte, 5)

	go proxyDestination(proxy.DestinationIP, proxy.DestinationPorts, chan_source, chan_destination)
	go proxySource(proxy.SourceIP, proxy.SourcePorts, chan_source, chan_destination)

	for {
		select {
		/*
			case nfo := <-proxyStream:
				fmt.Println("~~", nfo)*/
		case <-time.After(time.Second * 150):
			fmt.Println("sayonara")
			return nil
		}
	}
}

/* Create a ZMQ Proxy Reader from source of Proxy */
func proxyDestination(ip string, ports []int, chan_source chan []byte, chan_destination chan []byte) error {
	socket := ZmqRequestSocket(ip, ports)

	for {
		request := <-chan_destination
		reply, err_request := ZmqRequestByte(socket, request)
		if err_request != nil {
			fmt.Println("ERROR:", err_request)
			return err_request
		}
		chan_source <- reply
	}
}

/* Create a ZMQ Proxy Reader from source of Proxy */
func proxySource(ip string, ports []int, chan_source chan []byte, chan_destination chan []byte) error {
	socket := ZmqReplySocket(ip, ports)

	reply_handler := func(request []byte) []byte {
		chan_destination <- request
		reply := <-chan_source
		return reply
	}

	for {
		err_reply := ZmqReplyByte(socket, reply_handler)
		if err_reply != nil {
			fmt.Println("ERROR:", err_reply)
			return err_reply
		}
	}
	return nil
}
