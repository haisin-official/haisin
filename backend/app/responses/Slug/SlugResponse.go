package responses

type SlugGetResponse struct {
	Slug string `json:"slug" binding:"required"`
}

type SlugPostReponse struct {
	Slug string `json:"slug" binding:"required"`
}
