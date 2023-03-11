package usecases

import (
	"net/http"

	responses "github.com/haisin-official/haisin/app/http/responses/OAuth"
	"github.com/haisin-official/haisin/config"
)

func (OAuthUseCases) RedirectAction() (responses.OAuthRedirect, int, error) {
	url, code, err := config.GetRedirectUrl()

	if err != nil {
		return responses.OAuthRedirect{}, http.StatusInternalServerError, err
	}

	res := responses.OAuthRedirect{
		RedirectURL: url,
	}

	return res, code, err
}
