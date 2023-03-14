package usecases

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Service"
	responses "github.com/haisin-official/haisin/app/responses/Service"
	"github.com/haisin-official/haisin/app/utils"
	"github.com/haisin-official/haisin/database"
	"github.com/haisin-official/haisin/ent"
	"github.com/haisin-official/haisin/ent/service"
)

func (ServiceUseCase) ServicePostAction(req requests.ServicePostRequest) (responses.ServicePostResponse, int, error) {
	userId := req.UserID
	serviceName := req.Service
	serviceUrl := req.Url

	// URLが正しい形式かどうか確認
	if err := utils.CheckURL(serviceName, serviceUrl); err != nil {
		return responses.ServicePostResponse{}, http.StatusBadRequest, err
	}
	// URLをパース
	u, err := url.Parse(serviceUrl)
	if err != nil {
		return responses.ServicePostResponse{}, http.StatusInternalServerError, err
	}

	// クエリパラメータを削除した登録するURLを生成
	newUrl := &url.URL{
		Scheme: "https",
		Host:   u.Host,
		Path:   u.Path,
	}

	service, httpCode, err := register(userId, serviceName, newUrl.String())
	if err != nil {
		return responses.ServicePostResponse{}, httpCode, err
	}

	res := responses.ServicePostResponse{
		Service: service.Service.String(),
		Url:     service.URL,
	}
	return res, http.StatusOK, nil
}

func register(userId uuid.UUID, serviceName string, serviceUrl string) (*ent.Service, int, error) {
	client := database.GetClient()
	ctx := context.Background()

	// 存在するやつにする
	s := service.Service(serviceName)

	if err := service.ServiceValidator(s); err != nil {
		// サービスが存在しないのでエラーを返す
		return nil, http.StatusUnprocessableEntity, err
	}

	err := client.Service.
		Create().
		SetID(userId).
		SetService(s).
		OnConflictColumns(service.FieldService).
		UpdateNewValues().
		SetURL(serviceUrl).
		Exec(ctx)

	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	} else if ent.IsValidationError(err) {
		return nil, http.StatusUnprocessableEntity, err
	} else if ent.IsConstraintError(err) {
		return nil, http.StatusConflict, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("something error has occured \n%v", err)
	}

	// 新しく取得したサービスとURLを取得
	service, err := client.Service.
		Query().
		Unique(true).
		Select(service.FieldService, service.FieldURL).
		Where(service.IDEQ(userId)).
		Where(service.ServiceIn(s)).
		Only(ctx)

	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	} else if ent.IsValidationError(err) {
		return nil, http.StatusUnprocessableEntity, err
	} else if ent.IsConstraintError(err) {
		return nil, http.StatusConflict, err
	}
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("something error has occured \n%v", err)
	}

	return service, http.StatusOK, nil
}
