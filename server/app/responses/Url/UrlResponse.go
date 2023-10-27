package responses

type UrlSlugGetResponse struct {
	Slug     string                        `json:"slug" binding:"required,min=4,max=30"`
	Services *[]UrlSlugGetResponseServices `json:"services"`
}

type UrlSlugGetResponseServices struct {
	Service string `json:"service" binding:"required,service"`
	Url     string `json:"url" binding:"required,url"`
}

type UrlSlugServiceGetResponse struct {
	Url string `json:"url" binding:"required,url"`
}
