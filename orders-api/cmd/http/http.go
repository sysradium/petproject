package main

import (
	"log"
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

	server.Register()
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	server.Stop()
}
