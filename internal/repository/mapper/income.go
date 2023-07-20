package mapper

import (
	"github.com/personal-finance/domain/entity"
	"github.com/personal-finance/internal/repository/models"
)

func ToModelIncome(req *entity.Income) *models.Income {
	return &models.Income{
		UserId:            req.UserId,
		TotalIncome:       req.TotalIncome,
		IncomeInformation: req.IncomeInformation,
	}
}

func ToEntityIncome(req *models.Income) *entity.Income {
	return &entity.Income{
		ID:                req.ID,
		UserId:            req.UserId,
		TotalIncome:       req.TotalIncome,
		IncomeInformation: req.IncomeInformation,
		CreatedDate:       req.CreatedDate,
		UpdatedDate:       req.UpdatedDate,
	}
}

func ToArrEntityIncome(req []*models.Income) []*entity.Income {
	res := make([]*entity.Income, 0)

	for _, mod := range req {
		singRes := ToEntityIncome(mod)
		res = append(res, singRes)
	}

	return res
}
