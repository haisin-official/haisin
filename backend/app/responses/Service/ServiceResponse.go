package responses

type ServiceGetResponse struct {
	Slug     string                        `json:"slug" binding:"required,alphanum,min=4,max=30"`
	Services *[]ServiceGetResponseServices `json:"services"`
}

type ServiceGetResponseServices struct {
	Service string `json:"service" binding:"required,service"`
	Url     string `json:"url" binding:"required,url"`
}

type SericePostResponse struct {
	Service string `json:"service" binding:"required,service"`
	Url     string `json:"url" binding:"required,url"`
}

type ServiceDeleteResponse struct {
}
