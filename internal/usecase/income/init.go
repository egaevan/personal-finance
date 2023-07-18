package income

import (
	"github.com/personal-finance/internal/config"
	"github.com/personal-finance/internal/repository/mysql/income"
)

type incomeInteractor struct {
	cfg        *config.ServerConfig
	incomeRepo income.IncomeRepository
}

func NewIncomeInteractor(
	cfg *config.ServerConfig,
	incomeRepo income.IncomeRepository,
) *incomeInteractor {
	return &incomeInteractor{
		cfg:        cfg,
		incomeRepo: incomeRepo,
	}
}
