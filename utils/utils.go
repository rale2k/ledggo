package utils

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func GetSha256Hash(data string) string {
	x := sha256.Sum256([]byte(data))
	return hex.EncodeToString(x[:])
}
