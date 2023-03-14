package usecases

import (
	"net/http"
	"net/url"

	"github.com/google/uuid"
	requests "github.com/haisin-official/haisin/app/requests/Service"
	responses "github.com/haisin-official/haisin/app/responses/Service"
	"github.com/haisin-official/haisin/app/utils"
	"github.com/haisin-official/haisin/ent"
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

	resultUrl, httpCode, err := register(userId, serviceName, newUrl.String())
	if err != nil {
		return responses.ServicePostResponse{}, httpCode, err
	}

	res := responses.ServicePostResponse{
		Service: resultUrl.String(),
		Url:     resultUrl.URL,
	}
	return res, http.StatusOK, nil
}

func register(userId uuid.UUID, serviceName string, serviceUrl string) (*ent.Url, int, error) {
	return nil, http.StatusOK, nil
}
