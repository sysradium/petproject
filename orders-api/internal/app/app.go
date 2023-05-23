package app

import (
	"context"

	"github.com/sysradium/petproject/orders-api/internal/app/commands"
	"github.com/sysradium/petproject/orders-api/internal/app/queries"
	"github.com/sysradium/petproject/orders-api/internal/domain/order"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	BookOrder commands.BookOrderHandler
}

type Queries struct {
	ListBookedOrders queries.BookedOrdersHandler
}

type Publisher interface {
	Publish(ctx context.Context, event interface{}) error
}

func NewApplication(o order.Repository, e Publisher) App {
	return App{
		Commands: Commands{
			BookOrder: commands.NewBookOrderHandler(o, e),
		},
		Queries: Queries{
			ListBookedOrders: queries.NewBookedOrderHandler(o),
		},
	}
}
