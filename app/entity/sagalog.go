package entity

import (
	"encoding/json"
	"github.com/golang-module/carbon"
	"github.com/xuweijie4030/go-common/gosaga"
	"github.com/xuweijie4030/go-common/gosaga/proto"
	"gosagaapi/app/dao/db"
)

type SagaLog struct {
	Id        int         `gorm:"primaryKey" json:"id"`
	SagaId    int         `json:"saga_id"`
	SpanId    string      `json:"span_id"`
	Event     int         `json:"event"`
	Index     int         `json:"index"`
	Content   interface{} `json:"content"`
	CreatedAt string      `json:"created_at"`
}

func ConvToSagaLog(log db.SagaLog) (result SagaLog, err error) {
	result = SagaLog{
		Id:        log.Id,
		SagaId:    log.SagaId,
		SpanId:    log.SpanId,
		Event:     log.Event,
		Index:     log.Index,
		Content:   make(map[string]interface{}),
		CreatedAt: carbon.CreateFromTimestamp(log.CreatedAt).ToDateTimeString(),
	}

	switch log.Event {
	case gosaga.SagaCreatedEvent:
		result.Content = ""
	case gosaga.TransactionBeginEvent:
		result.Content = proto.Transaction{}
		if err = json.Unmarshal([]byte(log.Content), &result.Content); err != nil {
			return result, err
		}
	case gosaga.TransactionEndEvent:
		result.Content = make([]proto.TransactionResponse, 0)
		if err = json.Unmarshal([]byte(log.Content), &result.Content); err != nil {
			return result, err
		}
	case gosaga.TransactionAbortedEvent:
		result.Content = make([]proto.TransactionResponse, 0)
		if err = json.Unmarshal([]byte(log.Content), &result.Content); err != nil {
			return result, err
		}
	case gosaga.TransactionCompensateEvent:
		result.Content = make([]proto.CompensateResponse, 0)
		if err = json.Unmarshal([]byte(log.Content), &result.Content); err != nil {
			return result, err
		}
	case gosaga.SagaEndEvent:
		result.Content = ""
	}

	return result, nil
}

func MConvToSagaLog(logs []db.SagaLog) (result []SagaLog, err error) {
	result = make([]SagaLog, 0)
	for _, log := range logs {
		tempLog, err := ConvToSagaLog(log)
		if err != nil {
			return result, err
		}
		result = append(result, tempLog)
	}

	return result, nil
}
