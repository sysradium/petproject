package server

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/sysradium/petproject/orders-api/api"
	"github.com/sysradium/petproject/orders-api/internal/ports"
)

type Runner interface {
	Run(context.Context) error
}

type Server struct {
	e      *echo.Echo
	h      *ports.HttpServer
	router Runner
}

func (a *Server) Register() {
	api.RegisterHandlersWithBaseURL(
		a.e,
		a.h,
		"v1",
	)
}

func (a *Server) Start() error {
	go a.router.Run(context.Background())
	return a.e.Start(":8081")
}

func (a *Server) Stop() {
}

func New(e *echo.Echo, h *ports.HttpServer, router Runner) *Server {
	return &Server{
		e:      e,
		h:      h,
		router: router,
	}
}
