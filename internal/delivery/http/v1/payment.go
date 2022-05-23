package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/internal/usecase"
	"github.com/sirupsen/logrus"
)

type paymentRoutes struct {
	u usecase.Payment
}

func newPaymentRoutes(rg *gin.RouterGroup, u usecase.Payment) {
	r := &paymentRoutes{u}
	h := rg.Group("/payment")
	{
		h.GET("/", r.getPayment)
	}
}

type getPaymentRequest struct {
	Start time.Time `json:"start" binding:"required"`
	End   time.Time `json:"end" binding:"required"`
}

func (r *paymentRoutes) getPayment(ctx *gin.Context) {
	var req getPaymentRequest
	payments, err := r.u.GetPaymentHistory(ctx.Request.Context(), req.Start, req.End)
	if err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, payments)
}
