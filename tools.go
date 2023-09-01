//go:build tools
// +build tools

package sparetools

// https://github.com/google/wire/issues/299
import (
	_ "github.com/google/wire/cmd/wire"
)
