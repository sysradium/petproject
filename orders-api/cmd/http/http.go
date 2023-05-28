package main

import (
	"context"
	"log"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

func main() {
	server, cleanup, err := Initialize(
		"users-api:8080",
		"kafka:9092",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	shutdown := otelplay.ConfigureOpentelemetry(context.Background())
	defer shutdown()

	server.Register()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	server.Stop()
}
