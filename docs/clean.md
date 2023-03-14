### Routes
/routes
ginのルートを定義、パターンに合致したアクセスを受けるとControllersの関数を呼び出す

### Controllers
backend/app/http/controllers
ginから受け取ったリクエストを、UseCasesに伝える役割

### UseCases
backend/app/usecases
受け取ったリクエストを元に、データベースと接続
DBのレコードの結果からレスポンスを組み立て、Controllersに返す

### Requests / Responses
backend/app/http/requests
backend/app/http/responses

リクエストとレスポンスの構造体を定義

### Middleware
backend/app/http/middleware

* X-Requested-WithヘッダにXMLHttpRequestを持たない通信に404を返す
* 認証が必要なAPIにおいて、認証を行っていないユーザーに401を返す