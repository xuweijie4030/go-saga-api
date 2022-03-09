package agreement

import (
	"gosagaapi/app/entity"
)

type (
	GetSagaLogRequest struct {
		SagaId int `json:"saga_id" binding:"required"`
	}

	GetSagaLogResponse struct {
		Data []entity.SagaLog `json:"data"`
	}
)
