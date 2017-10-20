package golnats

import (
	"log"
	"time"

	"github.com/nats-io/go-nats"
)

/*
go-nats wrapper for quick pub-sub usage
*/

type Nats struct {
	Connection   *nats.Conn
	Subscription *nats.Subscription
	Channel      chan *nats.Msg
	Subject      string
	Message      []byte
	MsgReply     string
	Timeout      time.Duration
}

func PubSub(natsURL string, subject string) Nats {
	pubsub := Nats{}
	pubsub.Connect(natsURL)
	pubsub.Subject = subject
	return pubsub
}

func (natsClient *Nats) Log() {
	log.Println(string(natsClient.Message))
}

func (natsClient *Nats) Connect(natsURL string) (err error) {
	natsClient.Connection, err = nats.Connect(natsURL)
	if err != nil {
		log.Println("Connection to ", nats.DefaultURL, " failed. ", err.Error())
		return
	}
	log.Println("Connected to " + nats.DefaultURL)
	return
}

func (natsClient *Nats) Close() {
	natsClient.Connection.Close()
}

func (natsClient *Nats) Unsubscribe() {
	natsClient.Subscription.Unsubscribe()
}

func (natsClient *Nats) Request() (err error) {
	var msg *nats.Msg

	if natsClient.MsgReply == "" {
		natsClient.MsgReply = "gol-nats"
	}

	if natsClient.Timeout == (0 * time.Millisecond) {
		natsClient.Timeout = 10 * time.Millisecond
	}

	msg, err = natsClient.Connection.Request(natsClient.Subject, []byte(natsClient.MsgReply), natsClient.Timeout)
	if err != nil || msg == nil {
		log.Println(err)
		return
	}

	natsClient.Message = msg.Data
	return
}

func (natsClient *Nats) Reply() (err error) {
	natsClient.Connection.Subscribe(natsClient.Subject, func(m *nats.Msg) {
		natsClient.Connection.Publish(m.Reply, natsClient.Message)
	})
	return
}

func (natsClient *Nats) Publish() (err error) {
	natsClient.Connection.Publish(natsClient.Subject, natsClient.Message)
	return
}

func (natsClient *Nats) PublishMessage() (err error) {
	if natsClient.MsgReply == "" {
		natsClient.MsgReply = "gol-nats"
	}
	natsMsg := &nats.Msg{
		Subject: natsClient.Subject,
		Reply:   natsClient.MsgReply,
		Data:    natsClient.Message,
	}
	natsClient.Connection.PublishMsg(natsMsg)
	return
}

func (natsClient *Nats) SubscriberAsync() (err error) {
	natsClient.Connection.Subscribe(natsClient.Subject, func(msg *nats.Msg) {
		natsClient.Message = msg.Data
	})
	return
}

func (natsClient *Nats) SubscriberSync() (err error) {
	var msg *nats.Msg

	if natsClient.Subscription == nil {
		natsClient.Subscription, err = natsClient.Connection.SubscribeSync(natsClient.Subject)
		if err != nil {
			log.Println(err)
			return
		}
	}

	if natsClient.Timeout == (0 * time.Second) {
		natsClient.Timeout = 1 * time.Second
	}
	msg, err = natsClient.Subscription.NextMsg(natsClient.Timeout)
	if err != nil || msg == nil {
		log.Println(err)
		return
	}

	natsClient.Message = msg.Data
	return
}

func (natsClient *Nats) SubscriberChan() (err error) {
	var msg *nats.Msg

	if natsClient.Channel == nil {
		natsClient.Channel = make(chan *nats.Msg, 64)
	}
	if natsClient.Subscription == nil {
		natsClient.Subscription, err = natsClient.Connection.ChanSubscribe(natsClient.Subject, natsClient.Channel)
		if err != nil {
			log.Println(err)
			return
		}
	}

	msg = <-natsClient.Channel

	if msg == nil {
		log.Println(err)
		return
	}

	natsClient.Message = msg.Data
	return
}
