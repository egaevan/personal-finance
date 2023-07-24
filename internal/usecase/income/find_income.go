package income

import (
	"context"
	"github.com/personal-finance/domain/entity"
)

func (i *incomeInteractor) FindIncome(ctx context.Context, userId, id int) (*entity.Income, error) {
	res, err := i.incomeRepo.GetIncomeById(ctx, userId, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
