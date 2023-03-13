package validators

import (
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
	"github.com/haisin-official/haisin/app/utils"
)

type Ga struct {
	Ga string `validate:"ga"`
}

var gaValidator validator.Func = func(fl validator.FieldLevel) bool {
	// GA4の形式は G-XXXXXXXXXX である
	// 文字を取得
	ga4 := fl.Field().String()

	// 文字の長さを取得する
	// 12文字でなければエラーを返す
	if gaLen := utf8.RuneCountInString(ga4); gaLen != 12 {
		return false
	}

	// 初めの2文字を取得する
	// G-でなければエラーを返す
	if gaTwo := string([]rune(ga4)[:2]); gaTwo != "G-" {
		return false
	}
	// 後ろの10文字を取得する
	// 英数大文字であればtrueを返す
	// ひらがな,
	if ok := utils.CheckRegExp("[0-9A-Z]+", string([]rune(ga4)[2:12])); ok {
		return true
	}
	// 英数大文字でなければエラーを返す
	return false
}
