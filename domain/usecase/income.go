package usecase

import (
	"context"
	"github.com/personal-finance/domain/entity"
)

type IncomeUseCase interface {
	AddIncome(ctx context.Context, ett *entity.Income) error
	GetIncome(ctx context.Context, userId int) ([]*entity.Income, error)
}
