package income

import (
	"github.com/labstack/echo/v4"
	"github.com/personal-finance/domain/usecase"
)

type IncomeHandler struct {
	incomeUseCase usecase.IncomeUseCase
}

func NewHandler(e *echo.Echo, incomeUC usecase.IncomeUseCase) {
	handler := &IncomeHandler{
		incomeUseCase: incomeUC,
	}

	income := e.Group("/income")
	income.GET("/:userId", handler.GetIncome)
	income.GET("/:userId/:id", handler.FindIncome)
	income.POST("/:userId/add", handler.AddIncome)
	income.PUT("/:userId/:id", handler.PutIncome)
}
