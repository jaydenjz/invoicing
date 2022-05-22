package domain

import (
	"context"
	"time"
)

type Payment struct {
	Id            int64
	Amount        uint64    `json:"amount"`
	PaymentDate   time.Time `json:"paymentDate"`
	PaymentMethod string    `json:"paymentMethod"`
	CustomerName  string    `json:"customerName"`
}

type PaymentRepository interface {
	GetByDateRange(ctx context.Context, start, end time.Time) ([]Payment, error)
}
