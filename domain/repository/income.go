package repository

import (
	"context"
	"github.com/personal-finance/domain/entity"
)

type IncomeRepository interface {
	InsertIncome(ctx context.Context, ett *entity.Income) error
	GetIncome(ctx context.Context, userId int) ([]*entity.Income, error)
}
