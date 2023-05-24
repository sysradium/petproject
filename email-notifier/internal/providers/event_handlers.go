package providers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sysradium/petproject/email-notifier/internal/email"
	handler "github.com/sysradium/petproject/email-notifier/internal/handlers"
)

type EventHandlers map[string]func(msg *message.Message) error

func ProvideEventHandlers(
	notifier email.Sender,
) EventHandlers {
	return EventHandlers{
		"v1events.OrderBooked": handler.New(handler.NewOrderBookedHandler(notifier)),
	}
}
