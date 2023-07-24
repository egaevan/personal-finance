package income

import (
	"github.com/labstack/echo/v4"
	"github.com/personal-finance/pkg/utils"
	"net/http"
)

func (h *IncomeHandler) GetIncome(c echo.Context) error {

	ctx := c.Request().Context()
	reqId := c.Param("userId")
	userId := utils.ConvertStrToInt(reqId)

	res, err := h.incomeUseCase.GetIncome(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.RestBody{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})
}
