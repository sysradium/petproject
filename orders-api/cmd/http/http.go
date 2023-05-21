package main

import (
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sysradium/petproject/orders-api/api"
	"github.com/sysradium/petproject/orders-api/internal/handler"
	pbUsers "github.com/sysradium/petproject/users-api/proto/users/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	e := echo.New()
	e.HideBanner = true

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

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer conn.Close()

	api.RegisterHandlersWithBaseURL(
		e,
		handler.New(
			pbUsers.NewUsersServiceClient(conn),
		),
		"v1",
	)

	e.Logger.Fatal(e.Start(":8081"))
}
