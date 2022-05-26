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

var payments []domain.Payment

func New(pg *postgres.Postgres) *PaymentRepo {
	return &PaymentRepo{pg}
}

func (r *PaymentRepo) GetByDateRange(ctx context.Context, start, end time.Time) ([]domain.Payment, error) {
	//err := r.DB.Where("PaymentDate >= ? AND PaymentDate <= ?", start, end).Find(&payments).Error
	//err := r.DB.Debug().Table("Payments").Where("'PaymentDate' >= ? AND 'PaymentDate' <= ?", start, end).Find(&payments).Error
	err := r.DB.Table("Payments").Find(&payments).Error
	if err != nil {
		return []domain.Payment{}, err
	}
	return payments, nil
}
