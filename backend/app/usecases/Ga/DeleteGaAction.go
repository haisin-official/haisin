package usecases

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/http/requests/Ga"
	responses "github.com/haisin-official/haisin/app/http/responses/Ga"
	"github.com/haisin-official/haisin/database"
)

func (GaUseCases) DeleteGaAction(req requests.GaDeleteRequest) (responses.GaDeleteResponse, int, error) {
	userId := req.UserID

	httpCode, err := deleteGa(userId)
	if err != nil {
		return responses.GaDeleteResponse{}, httpCode, err
	}

	return responses.GaDeleteResponse{}, http.StatusOK, nil

}

func deleteGa(userId uuid.UUID) (int, error) {
	client := database.GetClient()
	ctx := context.Background()

	_, err := client.User.
		UpdateOneID(userId).
		SetNillableGa(nil).
		Save(ctx)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
