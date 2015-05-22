package golzmq

import "fmt"

/*
ZmqProxyConfig is structure to sustain ZMQ Proxy configuration.
*/
type ZmqProxyConfig struct {
	DestinationIP    string
	DestinationPorts []int
	SourceIP         string
	SourcePorts      []int
}

/*
ZmqProxySocket creates a Proxy connection for given ZmqProxyConfig.
*/
func ZmqProxySocket(proxy ZmqProxyConfig) {
	chanSource := make(chan []byte, 5)
	chanDestination := make(chan []byte, 5)

	go proxyDestination(proxy.DestinationIP, proxy.DestinationPorts, chanSource, chanDestination)
	go proxySource(proxy.SourceIP, proxy.SourcePorts, chanSource, chanDestination)
}

/*
proxyDestination creates a ZMQ Proxy Reader from source of Proxy.
*/
func proxyDestination(ip string, ports []int, chanSource chan []byte, chanDestination chan []byte) error {
	socket := ZmqRequestSocket(ip, ports)

	for {
		request := <-chanDestination
		reply, errRequest := ZmqRequestByte(socket, request)
		if errRequest != nil {
			fmt.Println("ERROR:", errRequest)
			return errRequest
		}
		chanSource <- reply
	}
}

/*
proxySource creates a ZMQ Proxy Reader from source of Proxy.
*/
func proxySource(ip string, ports []int, chanSource chan []byte, chanDestination chan []byte) error {
	socket := ZmqReplySocket(ip, ports)

	replyHandler := func(request []byte) []byte {
		chanDestination <- request
		reply := <-chanSource
		return reply
	}

	for {
		errReply := ZmqReplyByte(socket, replyHandler)
		if errReply != nil {
			fmt.Println("ERROR:", errReply)
			return errReply
		}
	}
}
