//go:build windows

// Package cmd is a package that contains subcommands for the spare CLI command.
package cmd

import (
	"os/exec"
)

func openBrowser(targetURL string) bool {
	return exec.Command("cmd", "/C", "start", "msedge", targetURL).Start() == nil
}
