package responses

type SlugGetResponse struct {
	Slug string `json:"slug" binding:"required,alphanumunicode,min=4,max=30"`
}

type SlugPostReponse struct {
	Slug string `json:"slug" binding:"required,alphanumunicode,min=4,max=30"`
}
