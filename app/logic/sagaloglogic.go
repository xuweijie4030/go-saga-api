package logic

import (
	"github.com/gin-gonic/gin"
	"gosagaapi/app/agreement"
	"gosagaapi/app/repository"
	"sync"
)

type SagaLogLogic struct {
}

var (
	sagaLogLogicOnce sync.Once
	sagaLogLogic     *SagaLogLogic
)

func NewSagaLogLogic() *SagaLogLogic {
	if sagaLogLogic != nil {
		return sagaLogLogic
	}
	sagaLogLogicOnce.Do(func() {
		sagaLogLogic = &SagaLogLogic{}
	})

	return sagaLogLogic
}

func (l *SagaLogLogic) GetSagaLog(ctx *gin.Context, sagaId int) (response agreement.GetSagaLogResponse, err error) {
	if response.Data, err = repository.NewSagaLogRepository().GetSagaLogBySagaId(ctx, sagaId); err != nil {
		return response, err
	}

	return response, nil
}
