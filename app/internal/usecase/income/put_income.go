package income

import (
	"context"
	"errors"
	"github.com/personal-finance/app/domain/entity"
)

func (i *incomeInteractor) PutIncomeById(ctx context.Context, req *entity.Income, userId, id int) error {
	res, err := i.incomeRepo.GetIncomeById(ctx, userId, id)
	if err != nil {
		return err
	}

	if res.UserId != userId && res.ID != id {
		return errors.New("not valid user")
	}

	err = i.incomeRepo.PutIncomeById(ctx, req, userId, id)
	if err != nil {
		return err
	}
	return nil
}
