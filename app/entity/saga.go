package entity

import (
	"encoding/json"
	"github.com/golang-module/carbon"
	"github.com/xuweijie4030/go-common/gosaga/proto"
	"gosagaapi/app/dao/db"
)

type Saga struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	TraceId   string         `json:"trace_id"`
	Content   proto.SagaInfo `json:"content"`
	Status    int            `json:"status"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

func ConvToSaga(log db.Saga) (result Saga, err error) {
	result = Saga{
		Id:        log.Id,
		TraceId:   log.TraceId,
		Status:    log.Status,
		CreatedAt: carbon.CreateFromTimestamp(log.CreatedAt).ToDateTimeString(),
		UpdatedAt: carbon.CreateFromTimestamp(log.UpdatedAt).ToDateTimeString(),
	}

	if err = json.Unmarshal([]byte(log.Content), &result.Content); err != nil {
		return result, err
	}

	return result, nil
}

func MConvToSaga(logs []db.Saga) (result []Saga, err error) {
	result = make([]Saga, 0)
	for _, log := range logs {
		tempLog, err := ConvToSaga(log)
		if err != nil {
			return result, err
		}
		result = append(result, tempLog)
	}

	return result, nil
}
