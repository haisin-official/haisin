package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
)

func GenRandToken() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		log.Panicf("Error with create new token 🚫\n %v", err)
	}

	return base64.StdEncoding.EncodeToString(b)
}
