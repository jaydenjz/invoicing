package httpserver

import (
	"net/http"
	"time"

	"github.com/jaydenjz/accounting/config"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*http.Server
}

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 10 * time.Second
)

// New -.
func New(handler http.Handler, config config.Config) error {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         config.Port,
	}

	logrus.Info("Service is running at: http://localhost", httpServer.Addr)
	return httpServer.ListenAndServe()
}
