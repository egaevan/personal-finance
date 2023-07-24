package income

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance/pkg/utils"
	"net/http"
)

func (h *incomeHandler) FindIncome(ctx *gin.Context) {

	reqId := ctx.Param("userId")
	id := ctx.Param("id")
	userId := utils.ConvertStrToInt(reqId)
	incomeId := utils.ConvertStrToInt(id)

	res, err := h.incomeUseCase.FindIncome(ctx, userId, incomeId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.RestBody{
			Message: "internal error",
		})

		return
	}

	ctx.JSON(http.StatusOK, utils.RestBody{
		Message: "success",
		Data:    res,
	})

	return
}
