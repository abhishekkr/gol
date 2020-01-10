package main

import (
	"fmt"
	"time"

	golnet "github.com/abhishekkr/gol/golnet"
)

func SomeReplies(bufRequest []byte) []byte {
	request := string(bufRequest)
	switch request {
	case "name":
		return []byte("some server")
	case "type":
		return []byte("golang TCP Listener")
	default:
		return []byte("something")
	}
}

func main() {
	go golnet.TCPServer("localhost:3456", SomeReplies)
	time.Sleep(time.Duration(2) * time.Second)
	client1 := golnet.CreateTCPClient("localhost:3456")
	client2 := golnet.CreateTCPClient("localhost:3456")
	defer client1.Connection.Close()
	defer client2.Connection.Close()

	name := client1.Request([]byte("name"))
	fmt.Println("name:", string(name))

	none := client2.Request([]byte("none"))
	fmt.Println("none:", string(none))
}
