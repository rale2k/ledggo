package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetSha256Hash(data string) string {
	x := sha256.Sum256([]byte(data))
	return hex.EncodeToString(x[:])
}
