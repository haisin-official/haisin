package responses

import "github.com/google/uuid"

type CallbackPostResponse struct {
	User CallbackPostUser `json:"user" binding:"required"`
}

type CallbackPostUser struct {
	Uuid  uuid.UUID `json:"uuid" binding:"required"`
	Email string    `json:"email" binding:"required"`
	Slug  string    `json:"slug" binding:"required"`
	Ga    *string   `json:"ga"`
}
