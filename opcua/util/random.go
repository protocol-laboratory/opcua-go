package util

import (
	"crypto/rand"
)

func GenerateRandomBytes(length int) []byte {
	bytes := make([]byte, length)
	// TODO better error handling
	rand.Read(bytes)
	return bytes
}
