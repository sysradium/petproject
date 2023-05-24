package handler

import (
	"fmt"

	events "github.com/sysradium/petproject/orders-api/api/events/v1"
)

type OrderBookedHandler struct {
}

func (o *OrderBookedHandler) NewEvent() interface{} {
	return &events.OrderBooked{}
}

func (o *OrderBookedHandler) Handle(e interface{}) error {
	msg := e.(*events.OrderBooked)
	fmt.Println(msg)

	return nil
}
