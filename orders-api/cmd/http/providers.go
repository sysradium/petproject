package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	events "github.com/sysradium/petproject/orders-api/api/events/v1"
	"github.com/sysradium/petproject/orders-api/internal/adapters/kafka/encoding"
)

type KafkaAddress string

type SendEmailOnOrderBooked struct {
}

func (o SendEmailOnOrderBooked) HandlerName() string {
	return "OnOrderBooked"
}

func (SendEmailOnOrderBooked) NewEvent() interface{} {
	return &events.OrderBooked{}
}

func (o SendEmailOnOrderBooked) Handle(ctx context.Context, e interface{}) error {
	event := e.(*events.OrderBooked)
	fmt.Println("received event", event)

	return nil
}

func newCQRSFacade(addr KafkaAddress) (*cqrs.Facade, *message.Router) {
	logger := watermill.NewStdLogger(false, false)
	cqrsMarshaler := encoding.ProtobufMarshaler{}

	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriber, err := kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{string(addr)},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         "test_consumer_group",
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{string(addr)},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic("unable to set up a publisher")
	}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	// router.AddMiddleware(middleware.Recoverer)

	cqrsFacade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		CommandsPublisher: publisher,
		CommandsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			return subscriber, nil
		},
		GenerateEventsTopic: func(eventName string) string {
			return "events"
		},
		EventsPublisher: publisher,
		Router:          router,
		GenerateCommandsTopic: func(commandName string) string {
			return commandName
		},
		EventHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
			return []cqrs.EventHandler{&SendEmailOnOrderBooked{}}
		},
		EventsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
			saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

			subscriber, err := kafka.NewSubscriber(
				kafka.SubscriberConfig{
					Brokers:               []string{string(addr)},
					Unmarshaler:           kafka.DefaultMarshaler{},
					OverwriteSaramaConfig: saramaSubscriberConfig,
					ConsumerGroup:         "test_consumer_group",
				},
				watermill.NewStdLogger(false, false),
			)
			return subscriber, err

		},
		CommandEventMarshaler: cqrsMarshaler,
		Logger:                logger,
	})

	if err != nil {
		log.Fatal(err)
	}

	return cqrsFacade, router
}

func main2() {
	facade, router := newCQRSFacade("localhost:9092")

	if err := facade.EventBus().Publish(context.Background(), &events.OrderBooked{Name: "vegetables"}); err != nil {
		log.Fatal(err)
	}

	if err := router.Run(context.Background()); err != nil {
		log.Fatal(err)
	}

}
