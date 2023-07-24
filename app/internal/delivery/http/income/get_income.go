package income

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance/pkg/utils"
	"net/http"
)

func (h *incomeHandler) GetIncome(ctx *gin.Context) {

	reqId := ctx.Param("userId")
	userId := utils.ConvertStrToInt(reqId)

	res, err := h.incomeUseCase.GetIncome(ctx, userId)
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
