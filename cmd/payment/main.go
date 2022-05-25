package main

import (
	"github.com/jaydenjz/accounting/config"
	"github.com/jaydenjz/accounting/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	config, err := config.New()
	if err != nil {
		logrus.Fatal(err)
	}

	app.Run(config)
}
