package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/internal/usecase"
	"github.com/jaydenjz/accounting/pkg/postgres"
	"github.com/spf13/viper"
)

func Run() {
	// Repository

	paymentUseCase := usecase.New()
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	// HTTP Server
	handler := gin.New()
	v1.newRouter(handler, paymentUseCase)
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func ReadConfigFile() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}
