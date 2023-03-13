# Validation
https://gin-gonic.com/ja/docs/examples/binding-and-validation/
Ginで標準使用されている https://github.com/go-playground/validator を使用する

# Validatorの基準
以下の基準を考える
## OAuth API
### Email
Validation, Emailを適用

## GA(Google Analytics) API
`G-XXXXXXXXXX`(Xは10個)の文字列
大文字英数字を必要とする

## Slug API
### Slug
4文字以上30文字以下の英数字