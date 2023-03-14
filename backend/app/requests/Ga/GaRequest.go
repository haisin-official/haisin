package requests

import "github.com/google/uuid"

type GaGetRequest struct {
	UserID uuid.UUID
}

type GaPostRequest struct {
	UserID uuid.UUID
	Ga     string
}

type GaPostRequestBody struct {
	Ga string `json:"ga" binding:"required,alphanumunicode,ga"`
}

type GaDeleteRequest struct {
	UserID uuid.UUID
}
