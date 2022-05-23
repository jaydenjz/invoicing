package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
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

func Run() {

	// HTTP Server
	testserver := &http.Server{
		Addr:         ":8081",
		Handler:      RouterTest(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g := new(errgroup.Group)
	g.Go(func() error {
		logrus.Info("App is running at http://localhost:8081")
		return testserver.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		logrus.Fatal(err)
	}
}
