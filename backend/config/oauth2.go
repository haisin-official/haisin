package config

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func GetOAuth2Conf() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     GetEnv("OAUTH_CLIENT"),
		ClientSecret: GetEnv("OAUTH_SECRET"),
		RedirectURL:  GetEnv("OAUTH_REDIRECT"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func GetRedirectUrl() (string, error) {
	conf := GetOAuth2Conf()

	url := conf.AuthCodeURL(GetEnv("OAUTH_CSRF_TOKEN"), oauth2.AccessTypeOffline)
	return url, nil
}

func CheckOAuthToken(state string, code string) (*oauth2.Token, error) {
	csrf := GetEnv("OAUTH_CSRF_TOKEN")
	// Tokenの値が違う場合, エラーを返却する
	if state != csrf {
		return nil, fmt.Errorf("CSRF TOKEN is invalid")
	}

	conf := GetOAuth2Conf()
	ctx := context.Background()

	token, err := conf.Exchange(ctx, code)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetUserInfo(token *oauth2.Token) (*v2.Userinfo, error) {
	conf := GetOAuth2Conf()
	ctx := context.Background()

	client := conf.Client(ctx, token)
	service, err := v2.NewService(ctx, option.WithHTTPClient(client))

	if err != nil {
		return nil, err
	}

	userInfo, err := service.Userinfo.Get().Context(ctx).Do()

	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
