package income

import (
	"github.com/labstack/echo/v4"
	"github.com/personal-finance/pkg/utils"
	"net/http"
)

func (h *IncomeHandler) FindIncome(c echo.Context) error {

	ctx := c.Request().Context()
	reqId := c.Param("userId")
	id := c.Param("id")
	userId := utils.ConvertStrToInt(reqId)
	incomeId := utils.ConvertStrToInt(id)

	res, err := h.incomeUseCase.FindIncome(ctx, userId, incomeId)
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
