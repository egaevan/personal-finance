package income

import (
	"github.com/gin-gonic/gin"
	"github.com/personal-finance/app/domain/entity"
	"github.com/personal-finance/pkg/utils"
	"net/http"
)

func (h *incomeHandler) AddIncome(ctx *gin.Context) {

	dataReq := entity.Income{}

	reqId := ctx.Param("userId")
	userId := utils.ConvertStrToInt(reqId)

	if err := ctx.Bind(&dataReq); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.RestBody{
			Message: "invalid data request",
		})
		return
	}

	dataReq.UserId = userId

	err := h.incomeUseCase.AddIncome(ctx, &dataReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.RestBody{
			Message: "internal error",
		})

		return
	}
	ctx.JSON(http.StatusCreated, utils.RestBody{
		Message: "success create",
	})
	return
}
