package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/internal/usecase"
)

func Run() {
	paymentUseCase := usecase.New(repository.)

	// HTTP Server
	handler := gin.Default()
	http.newRouter(handler, paymentUseCase)
	handler.Run(":5000")
}
