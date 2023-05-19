package main

import (
	"log"

	"github.com/ferminhg/be-dora-metrics/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
