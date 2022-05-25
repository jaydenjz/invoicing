package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/config"
	v1 "github.com/jaydenjz/accounting/internal/delivery/http/v1"
	"github.com/jaydenjz/accounting/internal/usecase"
	"github.com/jaydenjz/accounting/internal/usecase/repository"
	"github.com/jaydenjz/accounting/pkg/httpserver"
	"github.com/jaydenjz/accounting/pkg/postgres"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func RouterTest() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			"Welcome server 02",
		)
	})

	return e
}

func Run(cfg *config.Config) {
	// Postgres
	pg, err := postgres.New(cfg.Postgres.DatabaseUrl)
	if err != nil {
		logrus.Fatal("app - Run - postgres.New:", err)
	}
	defer pg.Close()

	// Use case
	paymentUseCase := usecase.New(repository.New(pg))

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, paymentUseCase)

	g.Go(func() error {
		return httpserver.New(handler)
	})

	if err := g.Wait(); err != nil {
		logrus.Fatal(err)
	}
}
