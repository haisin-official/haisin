package usecases

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Slug"
	responses "github.com/haisin-official/haisin/app/responses/Slug"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
)

func (SlugUseCase) SlugPostAction(req requests.SlugPostRequest) (responses.SlugPostReponse, int, error) {
	userId := req.UserID
	slug := req.Slug

	user, httpCode, err := changeSlug(userId, slug)
	if err != nil {
		return responses.SlugPostReponse{}, httpCode, err
	}

	res := responses.SlugPostReponse{
		Slug: user.Slug,
	}
	return res, http.StatusOK, nil
}

func changeSlug(userId uuid.UUID, slug string) (*ent.User, int, error) {
	client := database.GetClient()
	ctx := context.Background()

	// Transactionを発行して実行させるように変更
	// unique conflictの検証は後でやる
	user, err := client.User.
		UpdateOneID(userId).
		SetSlug(slug).
		Save(ctx)
	if ent.IsConstraintError(err) {
		return nil, http.StatusConflict, err
	} else if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil

}
