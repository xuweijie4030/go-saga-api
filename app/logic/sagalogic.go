package logic

import (
	"github.com/carefreex-io/logger"
	"github.com/gin-gonic/gin"
	"gosagaapi/app/agreement"
	"gosagaapi/app/repository"
	"sync"
)

type SagaLogic struct {
}

var (
	sagaLogic     *SagaLogic
	sagaLogicOnce sync.Once
)

func NewSagaLogic() *SagaLogic {
	if sagaLogic != nil {
		return sagaLogic
	}
	sagaLogicOnce.Do(func() {
		sagaLogic = &SagaLogic{}
	})

	return sagaLogic
}

func (l *SagaLogic) GetSaga(ctx *gin.Context, params agreement.GetSagaRequest) (response agreement.GetSagaResponse, err error) {
	if response.Total, err = repository.NewSagaRepository().CountSaga(ctx); err != nil {
		logger.ErrorfX(ctx, "repository.NewSagaRepository().CountPendingSaga failed: params=%v, err=%v", params, err)
		return response, err
	}

	skip := (params.Page - 1) * params.PageSize
	if response.Data, err = repository.NewSagaRepository().GetSaga(ctx, skip, params.PageSize); err != nil {
		logger.ErrorfX(ctx, "repository.NewSagaRepository().GetSaga failed: params=%v, err=%v", params, err)
		return response, err
	}

	return response, nil
}

func (l *SagaLogic) GetPendingSaga(ctx *gin.Context, params agreement.GetPendingSagaRequest) (response agreement.GetPendingSagaResponse, err error) {
	if response.Total, err = repository.NewSagaRepository().CountPendingSaga(ctx); err != nil {
		logger.ErrorfX(ctx, "repository.NewSagaRepository().CountPendingSaga failed: params=%v, err=%v", params, err)
		return response, err
	}

	skip := (params.Page - 1) * params.PageSize
	if response.Data, err = repository.NewSagaRepository().GetPendingSaga(ctx, skip, params.PageSize); err != nil {
		logger.ErrorfX(ctx, "repository.NewSagaRepository().GetPendingSaga failed: params=%v, err=%v", params, err)
		return response, err
	}

	return response, nil
}

func (l *SagaLogic) RecoverPendingSaga(ctx *gin.Context, params agreement.RecoverPendingSagaRequest) (response agreement.RecoverPendingSagaResponse, err error) {
	if err = repository.NewSagaRepository().RecoverPendingSaga(ctx, params.Id); err != nil {
		logger.ErrorfX(ctx, "repository.NewSagaRepository().RecoverPendingSaga failed: params=%v, err=%v", params, err)
		return response, err
	}

	return response, nil
}

func (l *SagaLogic) SubmitSaga(ctx *gin.Context, params agreement.SubmitSagaRequest) (response agreement.SubmitSagaResponse, err error) {
	if response.TraceId, err = repository.NewSagaRepository().SubmitSaga(ctx, params.SagaInfo); err != nil {
		logger.ErrorfX(ctx, "repository.NewSagaRepository().SubmitSaga failed: params=%v, err=%v", params, err)
		return response, err
	}

	return response, nil
}
