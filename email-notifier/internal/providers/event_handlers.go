package providers

import (
	"github.com/sysradium/petproject/email-notifier/internal/decorators"
	"github.com/sysradium/petproject/email-notifier/internal/email"
	"github.com/sysradium/petproject/email-notifier/internal/handlers"
	events "github.com/sysradium/petproject/orders-api/api/events/v1"
)

func ProvideEventHandlers(
	notifier email.Sender,
) handlers.EventHandlers {
	return handlers.EventHandlers{
		"v1events.OrderBooked": decorators.Unmarshal[*events.OrderBooked](
			handlers.NewOrderBookedHandler(notifier),
		),
	}
}
