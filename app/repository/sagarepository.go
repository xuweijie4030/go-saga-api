package repository

import (
	"github.com/carefreex-io/logger"
	"github.com/gin-gonic/gin"
	"github.com/xuweijie4030/go-common/gosaga/proto"
	"gosagaapi/app/dao/db"
	"gosagaapi/app/dao/rpc"
	"gosagaapi/app/entity"
	"gosagaapi/app/global"
	"sync"
)

type SagaRepository struct {
}

var (
	sagaRepository     *SagaRepository
	sagaRepositoryOnce sync.Once
)

func NewSagaRepository() *SagaRepository {
	if sagaRepository != nil {
		return sagaRepository
	}
	sagaRepositoryOnce.Do(func() {
		sagaRepository = &SagaRepository{}
	})

	return sagaRepository
}

func (r *SagaRepository) CountSaga(ctx *gin.Context) (result int64, err error) {
	if res := db.NewSagaDb(ctx).Query().Count(&result); res.Error != nil {
		logger.ErrorfX(ctx, "get saga count failed: err=%v", res.Error)
		return result, res.Error
	}

	return result, nil
}

func (r *SagaRepository) GetSaga(ctx *gin.Context, skip int, limit int) (result []entity.Saga, err error) {
	sagas := make([]db.Saga, 0)
	sagaDb := db.NewSagaDb(ctx)

	if res := sagaDb.Query().Where("id <= (?)", sagaDb.Query().Select("id").Order("id desc").Offset(skip).Limit(1)).Order("id desc").Limit(limit).Find(&sagas); res.Error != nil {
		logger.ErrorfX(ctx, "get saga failed: skip=%v, limit=%v, err=%v", skip, limit, res.Error)
		return result, res.Error
	}
	if result, err = entity.MConvToSaga(sagas); err != nil {
		logger.ErrorfX(ctx, "entity.MConvToSaga failed: sagas=%v, err=%v", sagas, err)
		return result, err
	}

	return result, nil
}

func (r *SagaRepository) CountPendingSaga(ctx *gin.Context) (result int64, err error) {
	if res := db.NewSagaDb(ctx).Query().Where("status = ?", global.SagaPendingStatus).Count(&result); res.Error != nil {
		logger.ErrorfX(ctx, "get pending saga count failed: err=%v", res.Error)
		return result, res.Error
	}

	return result, nil
}

func (r *SagaRepository) GetPendingSaga(ctx *gin.Context, skip int, limit int) (result []entity.Saga, err error) {
	sagas := make([]db.Saga, 0)
	sagaDb := db.NewSagaDb(ctx)

	if res := sagaDb.Query().Where("status = ?", global.SagaPendingStatus).Where("id <= (?)", sagaDb.Query().Select("id").Order("id desc").Offset(skip).Limit(1)).Order("id desc").Limit(limit).Find(&sagas); res.Error != nil {
		logger.ErrorfX(ctx, "get pending saga failed: skip=%v, limit=%v, err=%v", skip, limit, res.Error)
		return result, res.Error
	}
	if result, err = entity.MConvToSaga(sagas); err != nil {
		logger.ErrorfX(ctx, "entity.MConvToSaga failed: sagas=%v, err=%v", sagas, err)
		return result, err
	}

	return result, nil
}

func (r *SagaRepository) RecoverPendingSaga(ctx *gin.Context, ids []int) (err error) {
	if err = rpc.NewSagaRpc().RecoverPendingSaga(ctx, ids); err != nil {
		logger.ErrorfX(ctx, "rpc.NewSagaRpc().RecoverPendingSaga failed: ids=%v, err=%v", ids, err)
		return err
	}

	return nil
}

func (r *SagaRepository) SubmitSaga(ctx *gin.Context, sagaInfo proto.SagaInfo) (result string, err error) {
	if result, err = rpc.NewSagaRpc().SubmitSaga(ctx, sagaInfo); err != nil {
		logger.ErrorfX(ctx, "rpc.NewSagaRpc().SubmitSaga failed: sagaInfo=%v, err=%v", sagaInfo, err)
		return result, err
	}

	return result, nil
}
