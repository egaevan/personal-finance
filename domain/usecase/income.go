package usecase

import (
	"context"
	"github.com/personal-finance/domain/entity"
)

type IncomeUseCase interface {
	AddIncome(ctx context.Context, ett *entity.Income) error
}
