package util

import (
	"crypto/rand"
)

func GenerateRandomBytes(length int) []byte {
	bytes := make([]byte, length)
	// TODO better error handling
	_, _ = rand.Read(bytes)
	return bytes
}
