package rpc

import (
	"github.com/carefreex-io/logger"
	"github.com/gin-gonic/gin"
	"github.com/xuweijie4030/go-common/gosaga/proto"
	"github.com/xuweijie4030/go-common/gosaga/rpc"
	"sync"
)

type SagaRpc struct {
}

var (
	sagaRpc     *SagaRpc
	sagaRpcOnce sync.Once
)

func NewSagaRpc() *SagaRpc {
	if sagaRpc == nil {
		sagaRpcOnce.Do(func() {
			sagaRpc = &SagaRpc{}
		})
	}

	return sagaRpc
}

func (r *SagaRpc) RecoverPendingSaga(ctx *gin.Context, ids []int) (err error) {
	cli, err := rpc.NewClient()
	if err != nil {
		logger.ErrorfX(ctx, "rpc.NewClient failed: ids=%v, err=%v", ids, err)
		return err
	}

	request := proto.RecoverSagaRequest{
		SagaId: ids,
	}
	response := proto.RecoverSagaResponse{}
	if err = cli.RecoverSaga(ctx, &request, &response); err != nil {
		logger.ErrorfX(ctx, "cli.RecoverSaga failed: request=%v, err=%v", request, err)
		return err
	}

	return nil
}

func (r *SagaRpc) SubmitSaga(ctx *gin.Context, sagaInfo proto.SagaInfo) (result string, err error) {
	cli, err := rpc.NewClient()
	if err != nil {
		logger.ErrorfX(ctx, "rpc.NewClient failed: sagaInfo=%v, err=%v", sagaInfo, err)
		return result, err
	}

	request := sagaInfo
	response := proto.SubmitSagaResponse{}
	if err = cli.SubmitSaga(ctx, &request, &response); err != nil {
		logger.ErrorfX(ctx, "cli.SubmitSaga failed: request=%v, err=%v", request, err)
		return result, err
	}

	return response.TraceId, nil
}
