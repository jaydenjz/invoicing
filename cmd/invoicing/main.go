package main

import (
	"log"

	"github.com/jaydenjz/accounting/config"
	"github.com/jaydenjz/accounting/internal/app"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(config)
}
