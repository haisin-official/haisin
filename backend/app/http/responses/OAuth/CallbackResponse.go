package responses

import "github.com/google/uuid"

type OAuthCallback struct {
	User OAuthCallbackUser `json:"user" binding:"required"`
}

type OAuthCallbackUser struct {
	Uuid  uuid.UUID `json:"uuid" binding:"required"`
	Email string    `json:"email" binding:"required"`
	Slug  string    `json:"slug" binding:"required"`
	Ga    *string   `json:"ga"`
}
