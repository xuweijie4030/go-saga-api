package controller

import (
	"github.com/carefreex-io/xhttp"
	"github.com/gin-gonic/gin"
	"gosagaapi/app/agreement"
	"gosagaapi/app/logic"
	"sync"
)

type SagaLogController struct {
}

var (
	sagaLogController     *SagaController
	sagaLogControllerOnce sync.Once
)

func NewSagaLogController() *SagaController {
	if sagaLogController == nil {
		sagaLogControllerOnce.Do(func() {
			sagaLogController = &SagaController{}
		})
	}

	return sagaLogController
}

func (c *SagaController) GetSagaLog(ctx *gin.Context) {
	params := agreement.GetSagaLogRequest{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		xhttp.BadRequestResponse(ctx, err)
		return
	}

	response, err := logic.NewSagaLogLogic().GetSagaLog(ctx, params.SagaId)
	if err != nil {
		xhttp.InternalServerErrorResponse(ctx, err)
		return
	}

	xhttp.SuccessResponse(ctx, response)
}
