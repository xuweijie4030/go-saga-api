package validation

import (
	"github.com/carefreex-io/logger"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func RegisterValidation() {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		logger.Errorf("init validation failed")
		return
	}
	if err := v.RegisterValidation("demo", NewDemoValidation().Handle); err != nil {
		logger.Errorf("register demo validator failed: err=%v", err)
	}
}
