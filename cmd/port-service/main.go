package main

import (
	"log"
	"os"

	"github.com/evt/port-service/internal/config"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	// read config from env
	_ = config.Read()

	return nil
}
