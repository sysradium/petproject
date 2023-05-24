package handler

import (
	"fmt"

	"github.com/sysradium/petproject/email-notifier/internal/email"
	events "github.com/sysradium/petproject/orders-api/api/events/v1"
)

type OrderBookedHandler struct {
	notifier email.Sender
}

func (o *OrderBookedHandler) NewEvent() interface{} {
	return &events.OrderBooked{}
}

func (o *OrderBookedHandler) Handle(e interface{}) error {
	msg := e.(*events.OrderBooked)

	return o.notifier.Send(
		email.Message{
			Subject: fmt.Sprintf("An forder for %s placed", msg.Name),
		})
}

func NewOrderBookedHandler(n email.Sender) *OrderBookedHandler {
	return &OrderBookedHandler{
		notifier: n,
	}
}
