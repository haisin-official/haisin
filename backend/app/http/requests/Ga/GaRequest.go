package requests

import "github.com/google/uuid"

type GaGetRequest struct {
	UserID uuid.UUID
}

type GaPostRequest struct {
	UserID uuid.UUID
	Ga     string // new GA4 Code
}

type GaPostRequestBody struct {
	Ga string `json:"ga" binding:"required"`
}
