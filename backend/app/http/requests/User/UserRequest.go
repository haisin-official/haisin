package requests

import "github.com/google/uuid"

type UserGetRequest struct {
	UserID uuid.UUID
}
