package controller

import (
	"github.com/carefreex-io/logger"
	"github.com/carefreex-io/xhttp"
	"github.com/gin-gonic/gin"
	"gosagaapi/app/agreement"
	"gosagaapi/app/logic"
	"sync"
)

type SagaController struct {
}

var (
	sagaController     *SagaController
	sagaControllerOnce sync.Once
)

func NewSagaController() *SagaController {
	if sagaController == nil {
		sagaControllerOnce.Do(func() {
			sagaController = &SagaController{}
		})
	}

	return sagaController
}

func (c *SagaController) GetSaga(ctx *gin.Context) {
	params := agreement.GetSagaRequest{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		xhttp.BadRequestResponse(ctx, err)
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.PageSize == 0 {
		params.PageSize = 50
	}

	response, err := logic.NewSagaLogic().GetSaga(ctx, params)
	if err != nil {
		logger.ErrorfX(ctx, "logic.NewSagaLogic().GetSaga failed: params=%v err=%v", params, err)
		xhttp.InternalServerErrorResponse(ctx, err)
		return
	}

	xhttp.SuccessResponse(ctx, response)
}

func (c *SagaController) GetPendingSaga(ctx *gin.Context) {
	params := agreement.GetPendingSagaRequest{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		xhttp.BadRequestResponse(ctx, err)
		return
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.PageSize == 0 {
		params.PageSize = 50
	}

	response, err := logic.NewSagaLogic().GetPendingSaga(ctx, params)
	if err != nil {
		logger.ErrorfX(ctx, "logic.NewSagaLogic().GetPendingSaga failed: params=%v err=%v", params, err)
		xhttp.InternalServerErrorResponse(ctx, err)
		return
	}

	xhttp.SuccessResponse(ctx, response)
}

func (c *SagaController) RecoverPendingSaga(ctx *gin.Context) {
	params := agreement.RecoverPendingSagaRequest{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		xhttp.BadRequestResponse(ctx, err)
		return
	}

	response, err := logic.NewSagaLogic().RecoverPendingSaga(ctx, params)
	if err != nil {
		logger.ErrorfX(ctx, "logic.NewSagaLogic().RecoverPendingSaga failed: params=%v err=%v", params, err)
		xhttp.InternalServerErrorResponse(ctx, err)
		return
	}

	xhttp.SuccessResponse(ctx, response)
}

func (c *SagaController) SubmitSaga(ctx *gin.Context) {
	params := agreement.SubmitSagaRequest{}

	if err := ctx.ShouldBindJSON(&params); err != nil {
		xhttp.BadRequestResponse(ctx, err)
		return
	}

	response, err := logic.NewSagaLogic().SubmitSaga(ctx, params)
	if err != nil {
		logger.ErrorfX(ctx, "logic.NewSagaLogic().RecoverPendingSaga failed: params=%v err=%v", params, err)
		xhttp.InternalServerErrorResponse(ctx, err)
		return
	}

	xhttp.SuccessResponse(ctx, response)
}
