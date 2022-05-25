package repository

import (
	"context"
	"time"

	"github.com/jaydenjz/accounting/internal/domain"
	"github.com/jaydenjz/accounting/pkg/postgres"
)

type PaymentRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *PaymentRepo {
	return &PaymentRepo{pg}
}

func (r *PaymentRepo) GetByDateRange(ctx context.Context, start, end time.Time) ([]domain.Payment, error) {
	return []domain.Payment{}, nil
}
