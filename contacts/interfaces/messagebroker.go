package interfaces

import "context"

type Producer interface {
	Produce(context.Context) error
}

type Consumer interface {
	Consume(context.Context) error
}

// Broker
// Topic
// Message
