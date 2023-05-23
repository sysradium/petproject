package providers

import (
	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	waterllMiddleware "github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/sysradium/petproject/orders-api/internal/lib/common"
)

type KafkaAddress string

func NewRouter() (*message.Router, error) {
	logger := watermill.NewStdLogger(false, false)
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}
	router.AddMiddleware(waterllMiddleware.Recoverer)
	return router, nil
}

func NewCQRSFacade(
	addr KafkaAddress,
	router *message.Router,
	logger watermill.LoggerAdapter,
	handlers []cqrs.EventHandler,
) (*cqrs.EventBus, error) {
	cqrsMarshaler := common.ProtobufMarshaler{}

	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	// equivalent of auto.offset.reset: earliest
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest

	subscriberConfig := kafka.SubscriberConfig{
		Brokers:               []string{string(addr)},
		Unmarshaler:           kafka.DefaultMarshaler{},
		OverwriteSaramaConfig: saramaSubscriberConfig,
		ConsumerGroup:         "test_consumer_group",
	}
	subscriber, err := kafka.NewSubscriber(
		subscriberConfig,
		logger,
	)
	if err != nil {
		return nil, err
	}

	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{string(addr)},
			Marshaler: kafka.DefaultMarshaler{},
		},
		logger,
	)
	if err != nil {
		return nil, err
	}

	cqrsFacade, err := cqrs.NewFacade(cqrs.FacadeConfig{
		CommandsPublisher: publisher,
		CommandsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			return subscriber, nil
		},
		GenerateEventsTopic:   func(eventName string) string { return "events" },
		EventsPublisher:       publisher,
		Router:                router,
		GenerateCommandsTopic: func(commandName string) string { return commandName },
		EventHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.EventHandler {
			return handlers
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
				logger,
			)
			return subscriber, err

		},
		CommandEventMarshaler: cqrsMarshaler,
		Logger:                logger,
	})

	if err != nil {
		return nil, err
	}

	return cqrsFacade.EventBus(), nil
}
