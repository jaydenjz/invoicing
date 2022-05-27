package domain

import (
	"context"
	"time"
)

type Invoice struct {
	Id             int       `json:"id" gorm:"column:InvoiceId"`
	InvoiceNumber  int       `json:"invoiceNumber"`
	DueDate        time.Time `json:"dueDate"`
	BillingAddress string    `json:"billingAddress"`
	TotalAmount    float32   `json:"totalAmount"`
}

type InvoiceRepository interface {
	GetByDateRange(ctx context.Context, start, end time.Time) ([]Invoice, error)
}
