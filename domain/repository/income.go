package repository

import (
	"context"
	"github.com/personal-finance/domain/entity"
)

type IncomeRepository interface {
	InsertIncome(ctx context.Context, ett *entity.Income) error
	GetIncome(ctx context.Context, userId int) ([]*entity.Income, error)
	GetIncomeById(ctx context.Context, userId, id int) (*entity.Income, error)
	PutIncomeById(ctx context.Context, ett *entity.Income, userId, id int) error
}
