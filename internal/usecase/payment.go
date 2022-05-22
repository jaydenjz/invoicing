package usecase

import (
	"context"
	"time"

	"github.com/jaydenjz/accounting/internal/domain"
)

type (
	Payment interface {
		GetPaymentHistory(context.Context, time.Time, time.Time) ([]domain.Payment, error)
	}
)

type PaymentUseCase struct {
	repo domain.PaymentRepository
}

func New(r domain.PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{repo: r}
}

func (u *PaymentUseCase) GetPaymentHistory(ctx context.Context, start, end time.Time) (res []domain.Payment, err error) {
	payments, err := u.repo.GetByDateRange(ctx, start, end)
	if err != nil {
		return []domain.Payment{}, err
	}
	return payments, nil
}
