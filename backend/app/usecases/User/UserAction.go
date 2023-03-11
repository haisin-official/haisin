package usecases

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/http/requests/User"
	responses "github.com/haisin-official/haisin/app/http/responses/User"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/ent/user"
)

func (UserUseCases) UserGetAction(req requests.UserGetRequest) (responses.UserGetResponse, int, error) {
	// ユーザーIDを取得
	userId := req.UserID

	user, httpCode, err := getUser(userId)
	if err != nil {
		return responses.UserGetResponse{}, httpCode, err
	}

	// 返り値のデータを構築
	res := responses.UserGetResponse{
		User: responses.UserGetUser{
			Uuid:  user.ID,
			Email: user.Email,
			Slug:  user.Slug,
			Ga:    user.Ga,
		},
	}

	return res, http.StatusOK, nil

}

// UUIDからユーザー情報を取得
func getUser(userId uuid.UUID) (*ent.User, int, error) {
	client := database.GetClient()
	ctx := context.Background()
	res, err := client.User.
		Query().
		Unique(true).
		Select(user.FieldID, user.FieldGa, user.FieldEmail, user.FieldSlug).
		Where(user.IDEQ(userId)).
		Only(ctx)
	// ユーザーが見つからなかった場合
	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	}

	return res, http.StatusOK, nil
}
