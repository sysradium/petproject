//go:build wireinject
// +build wireinject

// wire.go
package main

import (
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sysradium/petproject/orders-api/api"
	"github.com/sysradium/petproject/orders-api/internal/app"
	"github.com/sysradium/petproject/orders-api/internal/handler"
	pbUsers "github.com/sysradium/petproject/users-api/proto/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newGrpcClient(addr string) (*grpc.ClientConn, func(), error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial(addr, opts...)
	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup, err
}

func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.DEBUG)

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = openapi3.Servers{
		{URL: "/v1"},
	}

	e.Use(
		echomiddleware.Logger(),
		middleware.OapiRequestValidator(swagger),
	)

	return e
}

func Initialize(addr string) (*app.App, func(), error) {
	wire.Build(
		newGrpcClient,
		wire.Bind(new(grpc.ClientConnInterface), new(*grpc.ClientConn)),
		pbUsers.NewUsersServiceClient,
		handler.New,
		NewEcho,
		app.New,
	)
	return &app.App{}, func() {}, nil
}
