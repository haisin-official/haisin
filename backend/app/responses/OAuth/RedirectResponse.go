package responses

type RedirectGetResponse struct {
	RedirectURL string `json:"redirect_url" binding:"required"`
}
