package commands

import (
	"context"

	"github.com/google/uuid"
	"github.com/sysradium/petproject/orders-api/internal/domain/order"
)

type BookOrderHandler CommandHandler[BookOrder, *order.Order]

type BookOrder struct {
	UserID uuid.UUID
	Name   string
}

type bookOrderHandler struct {
	repository order.Repository
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

	return newOrder, nil
}

func NewBookOrderHandler(o order.Repository) BookOrderHandler {
	return bookOrderHandler{
		repository: o,
	}
}
