package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xuweijie4030/go-common/gosaga"
	"github.com/xuweijie4030/go-common/gosaga/proto"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSagaController_SubmitSaga(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	data := proto.SubmitSagaRequest{
		ExecType:             gosaga.SerialExecType,
		//GlobalCompensateType: gosaga.BackwardCompensateType,
		Transaction: []proto.Transaction{
			{
				CallType:   5,
				Action:     "http://127.0.0.1:8888/transaction1",
				Compensate: "http://127.0.0.1:8888/compensate1",
				Data: map[string]interface{}{
					"test": "transaction1-test",
				},
				TransactionBaseOption: proto.TransactionBaseOption{
					//CompensateType: gosaga.BackwardCompensateType,
				},
			},
			{
				CallType:   gosaga.ApiTransaction,
				Action:     "http://127.0.0.1:8888/transaction2",
				Compensate: "http://127.0.0.1:8888/compensate2",
				Data: map[string]interface{}{
					"test": "transaction2-test",
				},
				TransactionBaseOption: proto.TransactionBaseOption{
					//CompensateType: gosaga.BackwardCompensateType,
				},
			},
		},
	}

	requestData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", "/sagas", bytes.NewReader(requestData))
	router.ServeHTTP(w, req)

	fmt.Println(w.Code)
	fmt.Println(w.Body.String())
}
