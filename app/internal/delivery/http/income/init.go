package income

import (
	"github.com/personal-finance/app/domain/usecase"
	"github.com/personal-finance/app/internal/config"
	incomeRepo "github.com/personal-finance/app/internal/repository/mysql/income"
	incomeUC "github.com/personal-finance/app/internal/usecase/income"
)

type incomeHandler struct {
	incomeUseCase usecase.IncomeUseCase
}

func NewHandler(cfg *config.ServerConfig, incomeRepo incomeRepo.IncomeRepository) *incomeHandler {
	incomeUseCase := incomeUC.NewIncomeInteractor(cfg, incomeRepo)
	return &incomeHandler{
		incomeUseCase: incomeUseCase,
	}
}
