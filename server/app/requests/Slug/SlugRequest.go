package requests

import "github.com/google/uuid"

type SlugGetRequest struct {
	UserID uuid.UUID
}

type SlugPostRequest struct {
	UserID uuid.UUID
	Slug   string
}

type SlugPostRequestBody struct {
	Slug string `json:"slug" binding:"required,alphanum,min=4,max=30"`
}
