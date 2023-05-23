package providers

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/sysradium/petproject/orders-api/internal/ports"
)

func NewEventHandlers() []cqrs.EventHandler {
	return []cqrs.EventHandler{
		&ports.SendEmailOnOrderBooked{},
	}
}
