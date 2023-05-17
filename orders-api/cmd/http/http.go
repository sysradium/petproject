package main

import (
	"fmt"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/sysradium/petproject/orders-api/api"
	"github.com/sysradium/petproject/orders-api/internal/handler"
)

func main() {
	e := echo.New()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(e, handler.New())

	e.Logger.Fatal(e.Start(":8080"))
}
