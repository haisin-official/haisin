package usecases

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Slug"
	responses "github.com/haisin-official/haisin/app/responses/Slug"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/ent/user"
)

func (SlugUseCases) SlugGetAction(req requests.SlugGetRequest) (responses.SlugGetResponse, int, error) {
	userId := req.UserID

	user, httpCode, err := getSlug(userId)
	if err != nil {
		return responses.SlugGetResponse{}, httpCode, err
	}

	res := responses.SlugGetResponse{
		Slug: user.Slug,
	}
	return res, http.StatusOK, nil
}

func getSlug(userId uuid.UUID) (*ent.User, int, error) {
	client := database.GetClient()
	ctx := context.Background()
	res, err := client.User.
		Query().Unique(true).Select(user.FieldSlug).
		Where(user.IDEQ((userId))).
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	}

	return res, http.StatusOK, nil
}
