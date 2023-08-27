// Package rand provides random values.
package rand

import (
	"crypto/rand"
	"math/big"
)

// RandomLowerAlphanumericStr returns a random string.
// String length is specified by length. String characters are a-z, 0-9.
func RandomLowerAlphanumericStr(length uint64) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		result[i] = charset[n.Int64()]
	}
	return string(result)
}
