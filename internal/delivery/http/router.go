package http

import (
	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/internal/usecase"
)

func newRouter(handler *gin.Engine, u usecase.Payment) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	v1 := handler.Group("/v1")
	{
		newPaymentRoutes(v1, u)
	}
}
