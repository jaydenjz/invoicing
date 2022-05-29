package domain

import (
	"context"
	"time"
)

func (Invoice) TableName() string {
	return "Invoices"
}

type Invoice struct {
	ID             int           `json:"id" gorm:"column:InvoiceId"`
	InvoiceNumber  int           `json:"invoiceNumber"  gorm:"column:InvoiceNumber"`
	DueDate        time.Time     `json:"dueDate" gorm:"column:DueDate"`
	BillingAddress string        `json:"billingAddress" gorm:"column:BillingAddress"`
	TotalAmount    float32       `json:"totalAmount" gorm:"column:TotalAmount"`
	Lines          []InvoiceLine `json:"lines" gorm:"foreignKey:InvoiceID"`
}

func (InvoiceLine) TableName() string {
	return "InvoiceLines"
}

type InvoiceLine struct {
	ID          int     `json:"id" gorm:"column:InvoiceLineId"`
	LineNo      int     `json:"lineNo" gorm:"column:SequenceNumber"`
	Name        string  `json:"name" gorm:"column:Name"`
	UnitPrice   float32 `json:"unitPrice" gorm:"column:UnitPrice"`
	Quantity    int     `json:"quantity" gorm:"column:Quantity"`
	TotalAmount float32 `json:"totalAmount" gorm:"column:TotalAmount"`
	InvoiceID   int     `json:"invoiceId" gorm:"column:InvoiceId"`
}

//go:generate mockgen -source=invoice.go -destination=../usecase/repository/invoice_test.go -package=repository_test
type InvoiceRepository interface {
	GetByDateRange(ctx context.Context, start, end time.Time) ([]Invoice, error)
}
