// Package xrand provides random values.
package xrand

import (
	"regexp"
	"testing"
)

func TestRandomAlphanumericStr(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		got, err := RandomLowerAlphanumericStr(10)
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != 10 {
			t.Errorf("RandomAlphanumericStr() = %v, want %v", got, 10)
		}

		if !isAlphanumeric(t, got) {
			t.Errorf("contains non-alphanumeric characters: %v", got)
		}
	})
}

func isAlphanumeric(t *testing.T, input string) bool {
	t.Helper()
	pattern := "^[a-z0-9]*$"
	match, _ := regexp.MatchString(pattern, input)
	return match
}
