package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		logrus.Debug("Service RUN on DEBUG mode")
	}
}

func main() {
	fmt.Println("Hello, World!")
}
