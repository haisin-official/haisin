package usecases

import (
	"context"
	"fmt"
	"net/http"

	responses "github.com/haisin-official/haisin/app/http/responses/OAuth"
	"github.com/haisin-official/haisin/app/utils"
	"github.com/haisin-official/haisin/config"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/ent/user"
)

func (OAuthUseCases) CallbackAction(state string, code string) (responses.OAuthCallback, int, error) {
	// 返り値のデータとHTTPコード
	res := responses.OAuthCallback{}
	httpCode := http.StatusOK

	// OAuthログインの検証
	token, httpCode, err := config.CheckOAuthToken(state, code)
	if err != nil {
		return res, httpCode, err
	}

	// OAuth2.0 を使用してEmailアドレスを取得
	userInfo, httpCode, err := config.GetUserInfo(token)

	if err != nil {
		return res, httpCode, err
	}

	userEmail := &userInfo.Email

	// Emailが取得できない場合, ユーザー登録ができないのでエラーを返却する
	if userEmail == nil {
		httpCode = http.StatusInternalServerError
		err = fmt.Errorf("could not get email address")
		return res, httpCode, err
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
			return res, httpCode, err
		}
	}
	// ユーザーデータを取得
	userData, httpCode, err := login(*userEmail)
	if err != nil {
		return res, httpCode, err
	}
	res.User.Uuid = userData.ID
	res.User.Slug = userData.Slug
	res.User.Email = userData.Email
	res.User.Ga = userData.Ga

	return res, httpCode, nil
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
