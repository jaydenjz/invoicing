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

var invoice domain.Invoice
var invoices []domain.Invoice

func New(pg *postgres.Postgres) *InvoiceRepo {
	return &InvoiceRepo{pg}
}

func (r *InvoiceRepo) GetByInvoiceNo(ctx context.Context, invcNo int) (*domain.Invoice, error) {
	err := r.DB.Model(domain.Invoice{InvoiceNumber: invcNo}).First(&invoice).Error
	if err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (r *InvoiceRepo) GetByDateRange(ctx context.Context, start, end time.Time) ([]domain.Invoice, error) {
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day()+1, 0, 0, 0, 0, start.Location()).Add(time.Duration(-1))
	err := r.DB.Debug().Preload("Lines").Find(&invoices, []int{1}).Error
	if err != nil {
		return []domain.Invoice{}, err
	}
	return invoices, nil
}
