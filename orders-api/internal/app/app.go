package app

import (
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

func NewApplication(o order.Repository) App {
	return App{
		Commands: Commands{
			BookOrder: commands.NewBookOrderHandler(o),
		},
		Queries: Queries{
			ListBookedOrders: queries.NewBookedOrderHandler(o),
		},
	}
}
