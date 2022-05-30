package usecase

import (
	"context"
	"time"

	"github.com/jaydenjz/accounting/internal/domain"
)

//go:generate mockgen -source=invoice_usecase.go -destination=./invoice_usecase_test.go -package=usecase_test
type Invoice interface {
	GetInvoiceByInvoiceNo(context.Context, int) (*domain.Invoice, error)
	GetInvoicesInDateRange(context.Context, time.Time, time.Time) ([]domain.Invoice, error)
}

type InvoiceUseCase struct {
	repo domain.InvoiceRepository
}

func New(r domain.InvoiceRepository) *InvoiceUseCase {
	return &InvoiceUseCase{repo: r}
}

func (u *InvoiceUseCase) GetInvoiceByInvoiceNo(ctx context.Context, invcNo int) (res *domain.Invoice, err error) {
	res, err = u.repo.GetByInvoiceNo(ctx, invcNo)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *InvoiceUseCase) GetInvoicesInDateRange(ctx context.Context, start, end time.Time) (res []domain.Invoice, err error) {
	res, err = u.repo.GetByDateRange(ctx, start, end)
	if err != nil {
		return []domain.Invoice{}, err
	}
	return res, nil
}
