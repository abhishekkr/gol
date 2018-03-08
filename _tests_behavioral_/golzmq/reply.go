package main

import (
	"fmt"

	golassert "github.com/abhishekkr/gol/golassert"
	golzmq "github.com/abhishekkr/gol/golzmq"
)

func get_n_put(messages []string) string {
	var ret_val string
	axn, msg := messages[0], messages[1]
	if axn == "get" {
		ret_val = fmt.Sprintf("GET: %s", msg)
	} else if axn == "put" {
		ret_val = fmt.Sprintf("PUT: %s", msg)
	} else {
		ret_val = fmt.Sprintf("unhandled request sent: %s", msg)
	}
	return ret_val
}

func ZmqReply() {
	socket := golzmq.ZmqReplySocket("127.0.0.1", []int{9797, 9898})
	for {
		err := golzmq.ZmqReply(socket, get_n_put)
		if err != nil {
			panic(err)
		}
	}
}

func ZmqRequest() {
	_socket := golzmq.ZmqRequestSocket("127.0.0.1", []int{9797, 9898})

	val, err := golzmq.ZmqRequest(_socket, "get", "anon")
	golassert.AssertEqual("GET: anon", val)
	golassert.AssertEqual(err, nil)

	val, err = golzmq.ZmqRequest(_socket, "put", "ymous")
	golassert.AssertEqual("PUT: ymous", val)
	golassert.AssertEqual(err, nil)
}

func main() {
	go ZmqReply()
	ZmqRequest()
	fmt.Println("passed not panic")
}
