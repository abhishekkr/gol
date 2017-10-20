package main

import (
	"runtime"

	golnats "github.com/abhishekkr/gol/golnats"
	nats "github.com/nats-io/go-nats"
)

/*
expects gnatsd server running in default mode

go get github.com/nats-io/nats
gnatsd
*/

func main() {
	pub := golnats.PubSub(nats.DefaultURL, "abk")
	pub.Message = []byte("whatever")

	sub := golnats.PubSub(nats.DefaultURL, "abk")

	go pub.PublishMessage()
	sub.Log()
	sub.Message = []byte("")
	sub.Log()
	sub.SubscriberSync()
	sub.Log()
	sub.Message = []byte("")
	sub.Log()
	sub.SubscriberSync()
	sub.Log()
	sub.Message = []byte("")
	sub.Log()
	sub.SubscriberSync()
	sub.Log()
	sub.Message = []byte("")
	sub.Log()

	runtime.Goexit()
}
