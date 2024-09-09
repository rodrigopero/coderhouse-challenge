package main

import (
	"github.com/rodrigopero/coderhouse-challenge/src/app"
	"log"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("error during app initialization: %s", err)
	}
}