package httpserver

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	*http.Server
}

const (
	_defaultReadTimeout  = 5 * time.Second
	_defaultWriteTimeout = 10 * time.Second
	_defaultAddr         = ":8081"
)

// New -.
func New(handler http.Handler) error {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
		Addr:         _defaultAddr,
	}

	logrus.Info("Service is running at: http://localhost", httpServer.Addr)
	return httpServer.ListenAndServe()
}
