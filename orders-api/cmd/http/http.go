package main

import (
	"log"
)

func main() {
	app, cleanup, err := Initialize("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	app.Register()
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	app.Stop()
}
