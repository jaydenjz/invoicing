package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jaydenjz/accounting/internal/domain"
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
		h.POST("/", r.addInvoice)
	}
}

type getInvoiceRequest struct {
	Start time.Time `json:"start" binding:"required"`
	End   time.Time `json:"end" binding:"required"`
}

func (r *InvoiceRoutes) getInvoiceByInvoiceNo(c *gin.Context) {
	invNo, err := strconv.Atoi(c.Param("invoiceNo"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Invalid param "+c.Param("invoiceNo"))
		return
	}
	invoices, err := r.service.GetInvoiceByInvoiceNo(c.Request.Context(), invNo)
	if err != nil {
		r.logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, invoices)
}

func (r *InvoiceRoutes) getInvoices(c *gin.Context) {
	//var req getPaymentRequest
	mockTime := time.Now()
	invoices, err := r.service.GetInvoicesInDateRange(c.Request.Context(), mockTime, mockTime)
	if err != nil {
		r.logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, invoices)
}

func (r *InvoiceRoutes) addInvoice(c *gin.Context) {
	var invoice *domain.Invoice
	if err := c.BindJSON(invoice); err != nil {
		r.logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	jsonStr, err := json.Marshal(invoice)
	if err != nil {
		r.logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	r.logger.Info(string(jsonStr))
	c.JSON(http.StatusOK, invoice)
}
