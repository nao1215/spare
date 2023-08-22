// Package main is verify-iap-receipt command entrypoint.
package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	t.Parallel()

	t.Run("Execute version subcommand", func(t *testing.T) {
		exitCode := -1
		oldOsExit := osExit
		osExit = func(code int) {
			exitCode = code
		}
		defer func() {
			osExit = oldOsExit
		}()

		os.Args = []string{
			"spare",
			"version",
		}
		main() // Run test

		if exitCode != 0 {
			t.Errorf("mismatch exit code: want=0, got=%d", exitCode)
		}
	})

	t.Run("Execution of a non-existent subcommand", func(t *testing.T) {
		exitCode := -1
		oldOsExit := osExit
		osExit = func(code int) {
			exitCode = code
		}
		defer func() {
			osExit = oldOsExit
		}()

		os.Args = []string{
			"spare",
			"non-existent",
		}
		main() // Run test

		if exitCode != 1 {
			t.Errorf("mismatch exit code: want=0, got=%d", exitCode)
		}
	})
}
