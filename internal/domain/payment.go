package domain

import (
	"context"
	"time"
)

type Payment struct {
	Id            int       `gorm:"column:PaymentId"`
	Amount        float32   `json:"amount"`
	PaymentDate   time.Time `json:"paymentDate"`
	PaymentMethod string    `json:"paymentMethod"`
	CustomerName  string    `json:"customerName"`
}

type PaymentRepository interface {
	GetByDateRange(ctx context.Context, start, end time.Time) ([]Payment, error)
}
