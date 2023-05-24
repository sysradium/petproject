package providers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	handler "github.com/sysradium/petproject/email-notifier/internal/handlers"
)

func ProvideHandlers(
	subscriber message.Subscriber,
	router *message.Router,
) error {
	router.AddNoPublisherHandler(
		"print_incoming_messages",
		"v1events.OrderBooked",
		subscriber,
		handler.New(&handler.OrderBookedHandler{}),
	)

	return nil
}
