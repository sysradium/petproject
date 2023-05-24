package app

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	handler "github.com/sysradium/petproject/email-notifier/internal/handlers"
)

type App struct {
	router     *message.Router
	subscriber message.Subscriber
}

func (a *App) Start() error {
	return a.router.Run(context.Background())
}

func (a *App) Stop() {}

func (a *App) RegisterHandlers() error {
	a.router.AddNoPublisherHandler(
		"print_incoming_messages",
		"v1events.OrderBooked",
		a.subscriber,
		handler.New(&handler.OrderBookedHandler{}),
	)

	return nil
}
func New(r *message.Router, s message.Subscriber) *App {
	return &App{
		router:     r,
		subscriber: s,
	}
}
