package agreement

import (
	"github.com/xuweijie4030/go-common/gosaga/proto"
	"gosagaapi/app/entity"
)

type (
	GetSagaRequest struct {
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	}

	GetSagaResponse struct {
		Total int64         `json:"total"`
		Data  []entity.Saga `json:"data"`
	}
)

type (
	GetPendingSagaRequest struct {
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	}

	GetPendingSagaResponse struct {
		Total int64         `json:"total"`
		Data  []entity.Saga `json:"data"`
	}
)

type (
	RecoverPendingSagaRequest struct {
		Id []int `json:"id" binding:"required"`
	}

	RecoverPendingSagaResponse struct {
	}
)

type (
	SubmitSagaRequest struct {
		proto.SagaInfo
	}

	SubmitSagaResponse struct {
		TraceId string `json:"trace_id"`
	}
)
