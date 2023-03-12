package usecases

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/http/requests/Ga"
	responses "github.com/haisin-official/haisin/app/http/responses/Ga"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/ent/user"
)

func (GaUseCases) GetGaAction(req requests.GaGetRequest) (responses.GaGetResponse, int, error) {
	userId := req.UserID

	user, httpCode, err := getGa(userId)
	if err != nil {
		return responses.GaGetResponse{}, httpCode, err
	}

	res := responses.GaGetResponse{
		Ga: user.Ga,
	}
	return res, http.StatusOK, nil
}

func getGa(userId uuid.UUID) (*ent.User, int, error) {
	client := database.GetClient()
	ctx := context.Background()
	res, err := client.User.
		Query().Unique(true).Select(user.FieldGa).
		Where(user.IDEQ((userId))).
		Only(ctx)
	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	}

	return res, http.StatusOK, nil
}
