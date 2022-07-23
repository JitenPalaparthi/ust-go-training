package messagebroker

import (
	"context"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type Messaging struct {
	Brokers []string
	Topic   string
	Message
	Config []interface{}
}

type Message struct {
	Key   []byte
	Value []byte
}

func (m *Messaging) Produce(ctx context.Context) error {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: m.Brokers,
		Topic:   m.Topic,
		// assign the logger to the writer
		Logger: l,
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: m.Key,
		// create an arbitrary message payload for the value
		Value: m.Value,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *Messaging) Consume(ctx context.Context) (*Message, error) {
	//l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: m.Brokers,
		Topic:   m.Topic,
		// assign the logger to the writer
		//Logger: l,
	})

	msg, err := r.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return &Message{Key: msg.Key, Value: msg.Value}, nil
}
