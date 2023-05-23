package queries

import (
	"context"

	"github.com/sysradium/petproject/orders-api/internal/domain/order"
)

type BookedOrders struct{}

type BookedOrdersHandler QueryHandler[BookedOrders, []*order.Order]

type bookedOrdersHandler struct {
	repo order.Repository
}

func (b bookedOrdersHandler) Handle(ctx context.Context, cmd BookedOrders) ([]*order.Order, error) {
	orders, err := b.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func NewBookedOrderHandler(r order.Repository) BookedOrdersHandler {
	return bookedOrdersHandler{
		repo: r,
	}
}
