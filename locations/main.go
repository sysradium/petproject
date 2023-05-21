package main

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
)

type Publisher interface {
	Publish(topic string, messages ...*message.Message) error
	Close() error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
	Close() error
}

func produce(ctx context.Context, publisher Publisher) {
	for range time.Tick(time.Second) {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))
		middleware.SetCorrelationID(watermill.NewUUID(), msg)
		err := publisher.Publish("example.topic", msg)
		if err != nil {
			panic(err)
		}

	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{"kafka:9092"},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(true, true),
	)
	if err != nil {
		panic(err)
	}

	produce(ctx, publisher)

}
