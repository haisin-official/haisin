package requests

import "github.com/google/uuid"

type ServiceGetRequest struct {
	UserID uuid.UUID
}

type ServicePostRequest struct {
	UserID  uuid.UUID
	Service string `json:"service" binding:"required,service"`
	Url     string `json:"url" binding:"required,url"`
}

type ServicePostRequestBody struct {
	Service string `json:"service" binding:"required,service"`
	Url     string `json:"url" binding:"required,url"`
}

type ServiceDeleteRequest struct {
	UserID  uuid.UUID
	Service string `json:"service" binding:"required,service"`
}

type ServiceDeleteRequestPath struct {
	Service string `json:"service" binding:"required,service"`
}
