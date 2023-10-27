package responses

import "github.com/google/uuid"

type UserGetResponse struct {
	User UserGetUser `json:"user" binding:"required"`
}

type UserGetUser struct {
	Uuid  uuid.UUID `json:"uuid" binding:"required"`
	Email string    `json:"email" binding:"required,email"`
	Slug  string    `json:"slug" binding:"required,alphanum,min=4,max=30"`
	Ga    *string   `json:"ga"`
}
