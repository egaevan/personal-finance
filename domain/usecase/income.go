package usecase

import (
	"context"
	"github.com/personal-finance/domain/entity"
)

type IncomeUseCase interface {
	AddIncome(ctx context.Context, req *entity.Income) error
	GetIncome(ctx context.Context, userId int) ([]*entity.Income, error)
	FindIncome(ctx context.Context, userId, id int) (*entity.Income, error)
	PutIncomeById(ctx context.Context, req *entity.Income, userId, id int) error
}
