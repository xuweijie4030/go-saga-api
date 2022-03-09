package validation

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

type DemoValidation struct {
}

var (
	demoValidation     *DemoValidation
	demoValidationOnce sync.Once
)

func NewDemoValidation() *DemoValidation {
	if demoValidation == nil {
		demoValidationOnce.Do(func() {
			demoValidation = &DemoValidation{}
		})
	}

	return demoValidation
}

func (v *DemoValidation) Handle(fl validator.FieldLevel) (result bool) {

	return true
}
