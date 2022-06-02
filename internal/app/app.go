package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/config"
	v1 "github.com/jaydenjz/accounting/internal/delivery/http/v1"
	"github.com/jaydenjz/accounting/internal/usecase"
	"github.com/jaydenjz/accounting/internal/usecase/repository"
	"github.com/jaydenjz/accounting/pkg/httpserver"
	"github.com/jaydenjz/accounting/pkg/postgres"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func Run(cfg *config.Config) {
	var logger *zap.Logger
	if cfg.Logger.Level == "info" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	// Postgres
	pg, err := postgres.New(cfg.Postgres.DatabaseUrl)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer pg.Close()

	// Use case
	invoiceUseCase := usecase.New(repository.New(pg))

	// HTTP Server
	router := gin.New()
	router.Use(cors.Default())
	v1.NewRouter(router, invoiceUseCase, logger)

	g.Go(func() error {
		return httpserver.New(router, *cfg)
	})

	if err := g.Wait(); err != nil {
		logger.Fatal(err.Error())
	}
}
