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
var lineItems []domain.InvoiceLine

func New(pg *postgres.Postgres) *InvoiceRepo {
	return &InvoiceRepo{pg}
}

func (r *InvoiceRepo) GetByDateRange(ctx context.Context, start, end time.Time) ([]domain.Invoice, error) {
	err := r.DB.Debug().Preload("Lines").Find(&invoices).Error
	if err != nil {
		return []domain.Invoice{}, err
	}
	return invoices, nil
}
