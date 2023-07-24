package income

import (
	"github.com/labstack/echo/v4"
	"github.com/personal-finance/domain/entity"
	"github.com/personal-finance/pkg/utils"
	"net/http"
)

func (h *IncomeHandler) PutIncome(c echo.Context) error {

	ctx := c.Request().Context()
	reqId := c.Param("userId")
	id := c.Param("id")
	userId := utils.ConvertStrToInt(reqId)
	incomeId := utils.ConvertStrToInt(id)

	dataReq := entity.Income{}

	if err := c.Bind(&dataReq); err != nil {
		c.JSON(http.StatusBadRequest, utils.RestBody{
			Message: "invalid data request",
		})
		return echo.ErrBadRequest
	}

	err := h.incomeUseCase.PutIncomeById(ctx, &dataReq, userId, incomeId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.RestBody{
			Message: "internal error",
		})

		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
	})
}
