package ports

import (
	"context"
	"fmt"

	events "github.com/sysradium/petproject/orders-api/api/events/v1"
)

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
