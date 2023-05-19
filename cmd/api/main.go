package main

import (
	"github.com/ferminhg/be-dora-metrics/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
