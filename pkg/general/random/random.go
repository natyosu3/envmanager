package random

import (
	"github.com/google/uuid"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func MakeRandomStringId() string {
	b := make([]byte, 64)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(b)
}

func MakeUuid() string {
	randomVal := uuid.New()
	
	return randomVal.String()
}

func MakeRandomNumberId() string {
	return uuid.New().String()[0:12]
}