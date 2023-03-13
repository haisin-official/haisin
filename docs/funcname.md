# リファクタリング
一度すべての機能を実装した後に、Go Linterを導入する


## ファイル名のルール
* `Path`:APIのパス. `Slug`, `User`, `Service` など...
* `Method`:APIのメソッド. `Get`, `Post`, `Delete` など...
### controllers
`PathController.go` の形式を取る. 例:`OAuthController.go`, `UserController.go`等

### UseCases
`PathUseCase.go` の形式を取る. 例:`UserUseCase.go`, `SlugUseCase.go`等

### Actions
`PathMethodAction.go` の形式を取る. 例:`UserGetAction.go`, `GaPostAction.go`等

### Requests / Responses
`PathRequest.go`,`PathResponse.go` の形式を取る. `SlugRequest.go`,`UserResponse.go`等

## 関数名・構造体のルール
### Controllers
`(PathController) PathMethod(c *gin.Context)`の形式を取る.
`(GaController) GaGet(c *gin.Context)`等

### Request / Response
構造体.  
`PathMethodRequest`,`PathMethodResponse`の形式を取る. リクエストボディは`PathMethodRequestBody`の形式を取る. `UserPostRequest`, `UserPostRequestBody`, `UserGetResponse`等

### Actions
`(PathUseCase) PathMethodAction (req requests.PathMethodRequest) (responses.PathMethodResponse, int, error)`の形式を取る.
`(GaUseCase) GaGetAction (req requests.GaGetRequest) (responses.GaGetResponse, int, error)`等