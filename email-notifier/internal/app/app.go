package app

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/sysradium/petproject/email-notifier/internal/handlers"
)

type App struct {
	router     *message.Router
	subscriber message.Subscriber
	handlers   handlers.EventHandlers
}

func (a *App) Start() error {
	return a.router.Run(context.Background())
}

func (a *App) Stop() {}

func (a *App) RegisterHandlers() error {
	for topic, handler := range a.handlers {
		a.router.AddNoPublisherHandler(
			"print_incoming_messages",
			topic,
			a.subscriber,
			handler,
		)
	}

	return nil
}
func New(
	r *message.Router,
	s message.Subscriber,
	handlers handlers.EventHandlers,
) *App {
	return &App{
		router:     r,
		subscriber: s,
		handlers:   handlers,
	}
}
