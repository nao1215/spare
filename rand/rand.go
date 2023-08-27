// Package rand provides random values.
package rand

import (
	"math/rand"
	"time"
)

// RandomLowerAlphanumericStr returns a random string.
// String length is specified by length. String characters are a-z, 0-9.
func RandomLowerAlphanumericStr(length uint64) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
