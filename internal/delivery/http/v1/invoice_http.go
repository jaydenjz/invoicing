package v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/internal/usecase"
	"github.com/sirupsen/logrus"
)

type InvoiceRoutes struct {
	u usecase.Invoice
}

func newInvoiceRoutes(rg *gin.RouterGroup, u usecase.Invoice) {
	r := &InvoiceRoutes{u}
	h := rg.Group("/invoice")
	{
		h.GET("/", r.getInvoice)
	}
}

type getInvoiceRequest struct {
	Start time.Time `json:"start" binding:"required"`
	End   time.Time `json:"end" binding:"required"`
}

func (r *InvoiceRoutes) getInvoice(ctx *gin.Context) {
	//var req getPaymentRequest
	mockTime := time.Now()
	payments, err := r.u.GetInvoices(ctx.Request.Context(), mockTime, mockTime)
	if err != nil {
		logrus.Error(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, payments)
}
