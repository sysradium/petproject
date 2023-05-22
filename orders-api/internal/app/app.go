package app

import (
	"github.com/labstack/echo/v4"
	"github.com/sysradium/petproject/orders-api/api"
	"github.com/sysradium/petproject/orders-api/internal/handler"
)

type App struct {
	e *echo.Echo
	h *handler.Handler
}

func (a *App) Register() {
	api.RegisterHandlersWithBaseURL(
		a.e,
		a.h,
		"v1",
	)
}

func (a *App) Start() error {
	return a.e.Start(":8081")
}

func (a *App) Stop() {
}

func New(e *echo.Echo, h *handler.Handler) *App {
	return &App{
		e: e,
		h: h,
	}
}
