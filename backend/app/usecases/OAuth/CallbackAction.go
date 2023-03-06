package usecases

import (
	"fmt"
	"net/http"

	responses "github.com/haisin-official/haisin/app/http/responses/OAuth"
	"github.com/haisin-official/haisin/config"
)

func (OAuthUseCases) CallbackAction(state string, code string) (responses.OAuthCallback, int, error) {
	res := responses.OAuthCallback{}

	token, httpCode, err := config.CheckOAuthToken(state, code)
	if err != nil {
		return res, httpCode, err
	}

	userInfo, httpCode, err := config.GetUserInfo(token)

	UserEmail := &userInfo.Email

	// Emailが取得できない場合, ユーザー登録ができないのでエラーを返却する
	if UserEmail == nil {
		return responses.OAuthCallback{}, http.StatusInternalServerError, fmt.Errorf("could not get email address")
	}

	if err != nil {
		return responses.OAuthCallback{}, httpCode, err
	}

	// ユーザー登録を開始もしくはログイン処理を行う
	// データベースから既存のemailが存在するか確認を行う

	return res, http.StatusOK, err
}
