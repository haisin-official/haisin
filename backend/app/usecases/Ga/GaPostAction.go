package usecases

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Ga"
	responses "github.com/haisin-official/haisin/app/responses/Ga"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
)

func (GaUseCase) GaPostAction(req requests.GaPostRequest) (responses.GaPostReponse, int, error) {
	userId := req.UserID
	ga := req.Ga

	user, httpCode, err := changeGa(userId, ga)
	if err != nil {
		return responses.GaPostReponse{}, httpCode, err
	}

	res := responses.GaPostReponse{
		Ga: *user.Ga,
	}
	return res, http.StatusOK, nil
}

func changeGa(userId uuid.UUID, ga string) (*ent.User, int, error) {
	client := database.GetClient()
	ctx := context.Background()

	// Transactionを発行して実行させるように変更
	// unique conflictの検証は後でやる
	user, err := client.User.
		UpdateOneID(userId).
		SetGa(ga).
		Save(ctx)
	if ent.IsConstraintError(err) {
		return nil, http.StatusConflict, err
	} else if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return user, http.StatusOK, nil

}
