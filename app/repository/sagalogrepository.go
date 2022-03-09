package repository

import (
	"github.com/carefreex-io/logger"
	"github.com/gin-gonic/gin"
	"gosagaapi/app/dao/db"
	"gosagaapi/app/entity"
	"sync"
)

type SagaLogRepository struct {
}

var (
	sagaLogRepositoryOnce sync.Once
	sagaLogRepository     *SagaLogRepository
)

func NewSagaLogRepository() *SagaLogRepository {
	if sagaLogRepository != nil {
		return sagaLogRepository
	}
	sagaLogRepositoryOnce.Do(func() {
		sagaLogRepository = &SagaLogRepository{}
	})

	return sagaLogRepository
}

func (r *SagaLogRepository) GetSagaLogBySagaId(ctx *gin.Context, sagaId int) (result []entity.SagaLog, err error) {
	sagaLogs := make([]db.SagaLog, 0)

	if res := db.NewSagaLogDb(ctx).DB.Where("saga_id = ?", sagaId).Find(&sagaLogs); res.Error != nil {
		logger.ErrorfX(ctx, "get saga log by saga id failed: sagaId=%v err=%v", sagaId, res.Error)
		return result, res.Error
	}

	if result, err = entity.MConvToSagaLog(sagaLogs); err != nil {
		logger.ErrorfX(ctx, "entity.MConvToSagaLog failed: sagaLogs=%v, err=%v", sagaLogs, err)
		return result, err
	}

	return result, nil
}
