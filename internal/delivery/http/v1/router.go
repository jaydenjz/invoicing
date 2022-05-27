package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/internal/usecase"
)

func NewRouter(handler *gin.Engine, u usecase.Invoice) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	v1 := handler.Group("/v1")
	{
		newInvoiceRoutes(v1, u)
	}
}
