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
	"github.com/haisin-official/haisin/ent/user"
)

func (ServiceUseCase) ServicePostAction(req requests.ServicePostRequest) (responses.ServicePostResponse, int, error) {
	userId := req.UserID
	name := req.Service
	redirect := req.Url

	// ServiceNameが存在するか確認
	if err := service.ServiceValidator(service.Service(name)); err != nil {
		return responses.ServicePostResponse{}, http.StatusBadRequest, err
	}

	// URLが正しい形式かどうか確認
	if err := utils.CheckURL(name, redirect); err != nil {
		return responses.ServicePostResponse{}, http.StatusBadRequest, err
	}
	// URLをパース
	u, err := url.Parse(redirect)
	if err != nil {
		return responses.ServicePostResponse{}, http.StatusInternalServerError, err
	}

	// クエリパラメータを削除した登録するURLを生成
	nu := &url.URL{
		Scheme: "https",
		Host:   u.Host,
		Path:   u.Path,
	}

	// userId & name ですでに登録されているデータが存在するか確認する
	service, _, _ := getService(userId, name)
	// 存在しない場合
	if service == nil {
		// insertを行う
		httpCode, err := register(userId, name, *nu)
		if err != nil {
			return responses.ServicePostResponse{}, httpCode, err
		}
	} else {
		// 存在する場合
		// updateを行う
		httpCode, err := update(service, *nu)
		if err != nil {
			return responses.ServicePostResponse{}, httpCode, err
		}
	}

	// 最新の情報を読み取る
	service, httpCode, err := getService(userId, name)
	if err != nil {
		return responses.ServicePostResponse{}, httpCode, err
	}

	res := responses.ServicePostResponse{
		Service: service.Service.String(),
		Url:     service.URL,
	}
	return res, http.StatusOK, nil
}

func getService(userId uuid.UUID, name string) (*ent.Service, int, error) {
	client := database.GetClient()
	ctx := context.Background()

	service, err := client.Service.
		Query().
		Select(service.FieldID, service.FieldService, service.FieldURL, service.UserIDColumn).
		Where(
			service.And(
				service.ServiceEQ(service.Service(name)),
				service.HasUserIDWith(user.IDEQ(userId)),
			),
		).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, http.StatusNotFound, err
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return service, http.StatusOK, nil
}

func register(userId uuid.UUID, name string, newURL url.URL) (int, error) {
	client := database.GetClient()
	ctx := context.Background()

	if err := service.ServiceValidator(service.Service(name)); err != nil {
		// サービスが存在しないのでエラーを返す
		fmt.Println("not found this service: ", name)
		return http.StatusUnprocessableEntity, err
	}

	// URLのIDを新しく生成
	u, err := utils.GenUUID()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// upsertはconflictに指定するカラムがuniqueの必要があるため使用できない。oauthの時と同様、丁寧にinsertすべきかupdateすべきか分岐させる必要がある
	if err := client.Service.
		Create().
		SetID(u).
		SetService(service.Service(name)).
		SetURL(newURL.String()).
		SetUserIDID(userId).
		Exec(ctx); ent.IsValidationError(err) {
		return http.StatusUnprocessableEntity, err
	} else if ent.IsConstraintError(err) {
		return http.StatusConflict, err
	} else if err != nil {
		return http.StatusInternalServerError, fmt.Errorf("something error has occured \n%v", err)
	}

	return http.StatusOK, nil

}

func update(service *ent.Service, newURL url.URL) (int, error) {
	client := database.GetClient()
	ctx := context.Background()

	if err := client.Service.
		UpdateOneID(service.ID).
		SetURL(newURL.String()).
		Exec(ctx); ent.IsNotFound(err) {
		return http.StatusNotFound, err
	} else if ent.IsValidationError(err) {
		return http.StatusUnprocessableEntity, err
	} else if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil

}
