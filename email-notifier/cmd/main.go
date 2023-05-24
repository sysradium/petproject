package main

import (
	"log"
)

func main() {

	a, err := Initialize()
	if err != nil {
		log.Fatal(err)
	}

	if err := a.RegisterHandlers(); err != nil {
		log.Fatal(err)
	}

	if err := a.Start(); err != nil {
		log.Fatal(err)
	}
}
