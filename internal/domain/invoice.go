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
	Lines          []InvoiceLine
}

type InvoiceLine struct {
	LineNo      int `gorm:"SequenceNumber"`
	Name        string
	UnitPrice   float32
	Quantity    int
	TotalAmount float32
	InvoiceId   int
	ItemId      int
}

//go:generate mockgen -source=invoice.go -destination=../usecase/repository/invoice_test.go -package=repository_test
type InvoiceRepository interface {
	GetByDateRange(ctx context.Context, start, end time.Time) ([]Invoice, error)
}
