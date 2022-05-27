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
	err := r.DB.Table("Invoices").Find(&invoices).Order("InvoiceNumber DESC").Error
	if err != nil {
		return []domain.Invoice{}, err
	}
	return invoices, nil
}
