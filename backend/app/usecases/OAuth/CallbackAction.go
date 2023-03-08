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
	res := responses.OAuthCallback{}
	httpCode := http.StatusOK

	token, httpCode, err := config.CheckOAuthToken(state, code)
	if err != nil {
		return res, httpCode, err
	}

	// Google OAuth を使用してEmailアドレスを取得
	userInfo, httpCode, err := config.GetUserInfo(token)

	if err != nil {
		return res, httpCode, err
	}

	UserEmail := &userInfo.Email

	// Emailが取得できない場合, ユーザー登録ができないのでエラーを返却する
	if UserEmail == nil {
		return res, http.StatusInternalServerError, fmt.Errorf("could not get email address")
	}

	// ユーザーデータを取得する
	client := database.GetClient()
	ctx := context.Background()
	userData, err := client.User.
		Query().
		Unique(true).
		Select(user.FieldID, user.FieldGa, user.FieldEmail, user.FieldSlug).
		Where(user.EmailEQ(*UserEmail)).
		Only(ctx)
	// ユーザーが存在しない場合
	if ent.IsNotFound(err) {
		res, httpCode, err := register(*UserEmail)
		if err != nil {
			return res, httpCode, err
		}
		return res, httpCode, err
	}
	// ユーザーが存在する場合
	res.User.Uuid = userData.ID
	res.User.Slug = userData.Slug
	res.User.Email = userData.Email
	res.User.Ga = userData.Ga

	return res, httpCode, nil
}

// 新規登録を行う
func register(userEmail string) (responses.OAuthCallback, int, error) {
	res := responses.OAuthCallback{}

	client := database.GetClient()
	ctx := context.Background()

	// 新しいUUIDを発行
	u, err := utils.GenUUID()
	if err != nil {
		return res, http.StatusInternalServerError, err
	}
	// 新しいSLUGを発行
	s, err := utils.GenSlug()
	if err != nil {
		return res, http.StatusInternalServerError, err
	}
	// UUIDとSLUGとEmailをデータベースに挿入
	err = client.User.
		Create().
		SetEmail(userEmail).
		SetSlug(s).
		SetID(u).
		Exec(ctx)
	if err != nil {
		return res, http.StatusInternalServerError, err
	}

	// レスポンスデータを構築する
	res.User.Email = userEmail
	res.User.Slug = s
	res.User.Uuid = u
	res.User.Ga = nil

	return res, http.StatusOK, nil
}
