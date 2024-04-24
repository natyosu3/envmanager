package random

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math/big"

	"github.com/google/uuid"
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
    randomBytes := make([]byte, 6)
    _, err := rand.Read(randomBytes)
    if err != nil {
		panic(err)
    }
    randomBigInt := new(big.Int).SetBytes(randomBytes)
    return fmt.Sprintf("%012d", randomBigInt)
}