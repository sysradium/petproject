package commands

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	events "github.com/sysradium/petproject/orders-api/api/events/v1"
	"github.com/sysradium/petproject/orders-api/internal/domain/order"
)

type BookOrderHandler CommandHandler[BookOrder, *order.Order]

type Publisher interface {
	Publish(ctx context.Context, event interface{}) error
}

type BookOrder struct {
	UserID uuid.UUID
	Name   string
}

type bookOrderHandler struct {
	repository order.Repository
	eventBus   Publisher
}

func (b bookOrderHandler) Handle(ctx context.Context, cmd BookOrder) (*order.Order, error) {
	newOrder, err := b.repository.Create(
		ctx,
		order.Order{
			Name:   cmd.Name,
			UserID: cmd.UserID,
		})

	if err != nil {
		return nil, err
	}

	if err := b.eventBus.Publish(ctx, &events.OrderBooked{
		Id:     newOrder.ID.String(),
		UserId: newOrder.UserID.String(),
		Name:   newOrder.Name,
	}); err != nil {
		fmt.Println(err)
	}

	return newOrder, nil
}

func NewBookOrderHandler(o order.Repository, e Publisher) BookOrderHandler {
	return bookOrderHandler{
		repository: o,
		eventBus:   e,
	}
}
