package main

import (
	"contacts/messagebroker"
	"context"
)

func main() {
	msg := new(messagebroker.Messaging)
	msg.Brokers = []string{"localhost:29092"}
	msg.Topic = "contacts-create"

	if msg, err := msg.Consume(context.Background()); err != nil {
		println(err.Error())
	} else {
		println(string(msg.Key), string(msg.Value))
	}
}
