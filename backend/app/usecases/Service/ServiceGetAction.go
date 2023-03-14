package usecases

import (
	"net/http"

	requests "github.com/haisin-official/haisin/app/requests/Service"
	responses "github.com/haisin-official/haisin/app/responses/Service"
)

func (ServiceUseCase) ServiceGetAction(req requests.ServiceGetRequest) (responses.ServiceGetResponse, int, error) {
	return responses.ServiceGetResponse{}, http.StatusOK, nil
}
