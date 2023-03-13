package validators

import (
	"github.com/go-playground/validator/v10"
)

type Ga struct {
	Ga string `validate:"ga"`
}

var gaValidator validator.Func = func(fl validator.FieldLevel) bool {

	return true
}
