package router

import (
	"github.com/gin-gonic/gin"
	"gosagaapi/app/controller"
)

func web(r *gin.Engine) {
	r.GET("/sagas", controller.NewSagaController().GetSaga)
	r.POST("/sagas", controller.NewSagaController().SubmitSaga)
	r.GET("/saga_log", controller.NewSagaLogController().GetSagaLog)
	r.GET("/pending_sagas", controller.NewSagaController().GetPendingSaga)
	r.POST("/recover_pending_saga", controller.NewSagaController().RecoverPendingSaga)
}
