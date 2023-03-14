package validators

import (
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

var regex = regexp.MustCompile("G-(([A-Z]|[0-9]){10})")

var gaValidator validator.Func = func(fl validator.FieldLevel) bool {
	// GA4の形式は G-XXXXXXXXXX である
	// 文字を取得
	ga := fl.Field().String()

	// 12文字, G-から始まる, 英数大文字であればtrueを返す
	return regex.MatchString(ga)
}
