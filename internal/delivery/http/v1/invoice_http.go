package v1

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/internal/usecase"
	"go.uber.org/zap"
)

type InvoiceRoutes struct {
	service usecase.Invoice
	logger  *zap.Logger
}

func newInvoiceRoutes(rg *gin.RouterGroup, u usecase.Invoice, logger *zap.Logger) {
	r := &InvoiceRoutes{u, logger}
	h := rg.Group("/invoice")
	{
		h.GET("/", r.getInvoices)
		h.GET("/:invoiceNo", r.getInvoiceByInvoiceNo)
	}
}

type getInvoiceRequest struct {
	Start time.Time `json:"start" binding:"required"`
	End   time.Time `json:"end" binding:"required"`
}

func (r *InvoiceRoutes) getInvoiceByInvoiceNo(ctx *gin.Context) {
	invNo, err := strconv.Atoi(ctx.Param("invoiceNo"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "Invalid param "+ctx.Param("invoiceNo"))
		return
	}
	invoices, err := r.service.GetInvoiceByInvoiceNo(ctx.Request.Context(), invNo)
	if err != nil {
		r.logger.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, invoices)
}

func (r *InvoiceRoutes) getInvoices(ctx *gin.Context) {
	//var req getPaymentRequest
	mockTime := time.Now()
	invoices, err := r.service.GetInvoicesInDateRange(ctx.Request.Context(), mockTime, mockTime)
	if err != nil {
		r.logger.Error(err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, invoices)
}
