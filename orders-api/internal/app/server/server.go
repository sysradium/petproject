package server

import (
	"github.com/labstack/echo/v4"
	"github.com/sysradium/petproject/orders-api/api"
	"github.com/sysradium/petproject/orders-api/internal/ports"
)

type Server struct {
	e *echo.Echo
	h *ports.HttpServer
}

func (a *Server) Register() {
	api.RegisterHandlersWithBaseURL(
		a.e,
		a.h,
		"v1",
	)
}

func (a *Server) Start() error {
	return a.e.Start(":8081")
}

func (a *Server) Stop() {
}

func New(e *echo.Echo, h *ports.HttpServer) *Server {
	return &Server{
		e: e,
		h: h,
	}
}
