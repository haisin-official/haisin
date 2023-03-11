package usecases

import (
	"context"
	"fmt"
	"net/http"

	requests "github.com/haisin-official/haisin/app/http/requests/OAuth"
	responses "github.com/haisin-official/haisin/app/http/responses/OAuth"
	"github.com/haisin-official/haisin/app/utils"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/ent/user"
)

func (OAuthUseCases) CallbackAction(req requests.CallbackRequest) (responses.OAuthCallback, int, error) {
	// stateとcodeを分解してOAuthログインの検証を行う
	token, httpCode, err := config.CheckOAuthToken(req.State, req.Code)
	if err != nil {
		return responses.OAuthCallback{}, httpCode, err
	}

	// OAuth2.0 を使用してEmailアドレスを取得
	userInfo, httpCode, err := config.GetUserInfo(token)
	if err != nil {
		return responses.OAuthCallback{}, httpCode, err
	}

	userEmail := &userInfo.Email

	// Emailが取得できない場合, ユーザー登録ができないのでエラーを返却する
	if userEmail == nil {
		httpCode = http.StatusInternalServerError
		err = fmt.Errorf("could not get email address")
		return responses.OAuthCallback{}, httpCode, err
	}

	// Emailが登録されているか確認を行う
	client := database.GetClient()
	ctx := context.Background()
	// userEmailの存在を確認する

	if exist := client.User.
		Query().
		Unique(true).
		Where(user.EmailEQ(*userEmail)).
		ExistX(ctx); !exist {
		// ユーザーが存在しない場合は新規登録する
		httpCode, err := register(*userEmail)
		if err != nil {
			return responses.OAuthCallback{}, httpCode, err
		}
	}
	// ユーザーデータを取得
	userData, httpCode, err := login(*userEmail)
	if err != nil {
		return responses.OAuthCallback{}, httpCode, err
	}

	// 返り値のデータを構築
	res := responses.OAuthCallback{
		User: responses.OAuthCallbackUser{
			Uuid:  userData.ID,
			Slug:  userData.Slug,
			Email: userData.Email,
			Ga:    userData.Ga,
		},
	}

	return res, http.StatusOK, nil
}

// 新規登録を行う
func register(userEmail string) (int, error) {
	// ここでは登録できたかどうかを確認するので、ユーザーデータは取得しない
	client := database.GetClient()
	ctx := context.Background()

	// 新しいUUIDを発行
	u, err := utils.GenUUID()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// 新しいSLUGを発行
	s, err := utils.GenSlug()
	if err != nil {
		return http.StatusInternalServerError, err
	}
	// UUIDとSLUGとEmailをデータベースに挿入
	err = client.User.
		Create().
		SetEmail(userEmail).
		SetSlug(s).
		SetID(u).
		Exec(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func login(userEmail string) (*ent.User, int, error) {
	client := database.GetClient()
	ctx := context.Background()

	res, err := client.User.
		Query().
		Unique(true).
		Select(user.FieldID, user.FieldGa, user.FieldEmail, user.FieldSlug).
		Where(user.EmailEQ(userEmail)).
		Only(ctx)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return res, http.StatusOK, nil
}
