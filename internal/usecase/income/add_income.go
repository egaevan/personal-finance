package income

import (
	"context"
	"github.com/personal-finance/domain/entity"
)

func (i *incomeInteractor) AddIncome(ctx context.Context, req *entity.Income) error {
	err := i.incomeRepo.InsertIncome(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
