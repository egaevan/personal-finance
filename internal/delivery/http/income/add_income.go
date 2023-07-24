package income

import (
	"github.com/labstack/echo/v4"
	"github.com/personal-finance/domain/entity"
	"github.com/personal-finance/pkg/utils"
	"net/http"
)

func (h *IncomeHandler) AddIncome(c echo.Context) error {

	ctx := c.Request().Context()
	dataReq := entity.Income{}

	reqId := c.Param("userId")
	userId := utils.ConvertStrToInt(reqId)

	if err := c.Bind(&dataReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.RestBody{
			Message: "invalid data request",
		})
		return echo.ErrBadRequest
	}

	dataReq.UserId = userId

	err := h.incomeUseCase.AddIncome(ctx, &dataReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.RestBody{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}
	return c.JSON(http.StatusCreated, utils.RestBody{
		Message: "success create",
	})
}
