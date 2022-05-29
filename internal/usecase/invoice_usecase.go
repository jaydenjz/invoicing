package usecase

import (
	"context"
	"time"

	"github.com/jaydenjz/accounting/internal/domain"
)

//go:generate mockgen -source=invoice_usecase.go -destination=./invoice_usecase_test.go -package=usecase_test
type (
	Invoice interface {
		GetInvoices(context.Context, time.Time, time.Time) ([]domain.Invoice, error)
	}
)

type InvoiceUseCase struct {
	repo domain.InvoiceRepository
}

func New(r domain.InvoiceRepository) *InvoiceUseCase {
	return &InvoiceUseCase{repo: r}
}

func (u *InvoiceUseCase) GetInvoices(ctx context.Context, start, end time.Time) (res []domain.Invoice, err error) {
	payments, err := u.repo.GetByDateRange(ctx, start, end)
	if err != nil {
		return []domain.Invoice{}, err
	}
	return payments, nil
}
