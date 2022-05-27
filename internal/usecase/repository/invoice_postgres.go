package repository

import (
	"context"
	"time"

	"github.com/jaydenjz/accounting/internal/domain"
	"github.com/jaydenjz/accounting/pkg/postgres"
)

type InvoiceRepo struct {
	*postgres.Postgres
}

var invoices []domain.Invoice

func New(pg *postgres.Postgres) *InvoiceRepo {
	return &InvoiceRepo{pg}
}

func (r *InvoiceRepo) GetByDateRange(ctx context.Context, start, end time.Time) ([]domain.Invoice, error) {
	//err := r.DB.Where("PaymentDate >= ? AND PaymentDate <= ?", start, end).Find(&payments).Error
	//err := r.DB.Debug().Table("Payments").Where("'PaymentDate' >= ? AND 'PaymentDate' <= ?", start, end).Find(&payments).Error
	err := r.DB.Table("Invoices").Find(&invoices).Error
	if err != nil {
		return []domain.Invoice{}, err
	}
	return invoices, nil
}
