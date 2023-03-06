package responses

type OAuthRedirect struct {
	RedirectURL string `json:"redirect_url" binding:"required"`
}
