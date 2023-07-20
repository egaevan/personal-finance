package income

import (
	"github.com/labstack/echo/v4"
	"github.com/personal-finance/domain/entity"
	"net/http"
)

type responseError struct {
	Message string `json:"message"`
}

func (h *IncomeHandler) AddIncome(c echo.Context) error {

	ctx := c.Request().Context()
	dataReq := entity.Income{}

	if err := c.Bind(&dataReq); err != nil {
		c.JSON(http.StatusBadRequest, responseError{
			Message: "invalid data request",
		})
		return echo.ErrBadRequest
	}

	err := h.incomeUseCase.AddIncome(ctx, &dataReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responseError{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, responseError{
		Message: "success create",
	})
}
