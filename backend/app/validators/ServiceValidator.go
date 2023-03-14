package validators

import "github.com/go-playground/validator/v10"

var serviceList = map[string]bool{
	"twitter": true,
	"youtube": true,
}

var serviceValidator validator.Func = func(fl validator.FieldLevel) bool {
	if service, ok := fl.Field().Interface().(string); ok {
		return serviceList[service]
	}
	return false
}
