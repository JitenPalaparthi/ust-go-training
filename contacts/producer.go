package main

import (
	"contacts/messagebroker"
	"context"
)

func main() {
	msg := new(messagebroker.Messaging)
	msg.Brokers = []string{"localhost:29092"}
	msg.Topic = "contacts-create"
	msg.Key = []byte("contacts")
	msg.Value = []byte("Contact Jiten, male, bangalore")
	if err := msg.Produce(context.Background()); err != nil {
		println(err.Error())
	} else {
		println("message successfuly produced")
	}
}
