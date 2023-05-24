//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/sysradium/petproject/email-notifier/internal/app"
	"github.com/sysradium/petproject/email-notifier/internal/providers"
)

func Initialize() (*app.App, error) {
	wire.Build(
		providers.ProvideLogger,
		providers.ProvideSubscriber,
		providers.ProvideRouter,
		providers.ProvideEventHandlers,
		providers.ProvideEmailNotifier,
		app.New,
	)
	return &app.App{}, nil
}
