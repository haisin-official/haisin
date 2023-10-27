package utils

import "github.com/google/uuid"

func GenUUID() (uuid.UUID, error) {
	u, err := uuid.NewRandom()

	if err != nil {
		return u, err
	}

	return u, nil
}
