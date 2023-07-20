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
