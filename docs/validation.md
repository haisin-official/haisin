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

### Service Provider
* twitter or youtube
* これに関してはCustomValidatorを実装して良さそう
* https://www.miracleave.co.jp/contents/1354/gin-make-validator/
### Service URL
* httpsで始まっているか
* 2083文字以下か
---
* twitter or youtube
* ドメインがtwitter.com or youtube.com
* 悪いパラメーターによって強制リダイレクトされないよう、Queryは削除する
* Serviceのセットに関しては文字列の操作や置き換えが入ると思うので、Validatorとしてルールを作成はしない(ただし、ドメインが合っているか、クエリを取り除いた正しいURLはなにか等、検証する必要はある)
* それ以外のValidation(文字が必要とか、長すぎてもだめとか)そういうものは入る
* net/urlパッケージを使用してHost部を抜き出す事で、ドメインだけを抜き出す事が可能