package usecases

import (
	"context"
	"net/http"

	requests "github.com/haisin-official/haisin/app/requests/Service"
	responses "github.com/haisin-official/haisin/app/responses/Service"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
)

func (ServiceUseCase) ServiceDeleteAction(req requests.ServiceDeleteRequest) (responses.ServiceDeleteResponse, int, error) {
	userId := req.UserID
	name := req.Service

	// 指定されたデータがあるか確認
	service, httpCode, err := getService(userId, name)
	if err != nil {
		return responses.ServiceDeleteResponse{}, httpCode, err
	}
	// 存在すれば削除
	httpCode, err = deleteService(service)
	if err != nil {
		return responses.ServiceDeleteResponse{}, httpCode, err
	}

	return responses.ServiceDeleteResponse{}, http.StatusOK, nil
}

func deleteService(service *ent.Service) (int, error) {
	client := database.GetClient()
	ctx := context.Background()
	// userIDのUUIDとNameを条件にかけてDROPする
	if err := client.Service.
		DeleteOne(service).
		Exec(ctx); ent.IsNotFound(err) {
		return http.StatusNotFound, err
	} else if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
