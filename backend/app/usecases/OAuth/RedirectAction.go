package usecases

import (
	"net/http"

	requests "github.com/haisin-official/haisin/app/requests/OAuth"
	responses "github.com/haisin-official/haisin/app/responses/OAuth"
	"github.com/haisin-official/haisin/config"
)

func (OAuthUseCase) RedirectGetAction(req requests.RedirectGetRequest) (responses.RedirectGetResponse, int, error) {
	url, code, err := config.GetRedirectUrl()

	if err != nil {
		return responses.RedirectGetResponse{}, http.StatusInternalServerError, err
	}

	res := responses.RedirectGetResponse{
		RedirectURL: url,
	}

	return res, code, err
}
