package requests

type UrlSlugGetRequest struct {
	Slug string `json:"slug" binding:"required,min=4,max=30"`
}

type UrlSlugServiceGetRequest struct {
	Slug    string `json:"slug" binding:"required,min=4,max=30"`
	Service string `json:"service" binding:"required,service"`
}
