package config

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func GetOAuth2Conf() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT"),
		ClientSecret: os.Getenv("OAUTH_SECRET"),
		RedirectURL:  os.Getenv("OAUTH_REDIRECT"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func GetRedirectUrl() (string, int, error) {
	conf := GetOAuth2Conf()
	url := conf.AuthCodeURL(os.Getenv("OAUTH_CSRF_TOKEN"), oauth2.AccessTypeOffline)
	return url, http.StatusOK, nil
}

func CheckOAuthToken(state string, code string) (*oauth2.Token, int, error) {
	csrf := os.Getenv("OAUTH_CSRF_TOKEN")
	// Tokenの値が違う場合, エラーを返却する
	if state != csrf {
		return nil, http.StatusForbidden, fmt.Errorf("CSRF TOKEN Invalid")
	}

	conf := GetOAuth2Conf()
	ctx := context.Background()

	token, err := conf.Exchange(ctx, code)

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return token, http.StatusOK, nil
}

func GetUserInfo(token *oauth2.Token) (*v2.Userinfo, int, error) {
	conf := GetOAuth2Conf()
	ctx := context.Background()

	client := conf.Client(ctx, token)
	service, err := v2.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	userInfo, err := service.Userinfo.Get().Context(ctx).Do()

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return userInfo, http.StatusOK, nil
}
