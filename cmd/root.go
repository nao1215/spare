// Package cmd is a package that contains subcommands for the spare CLI command.
package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// newRootCmd creates a new root command. This command is the entry point of the CLI.
// It is responsible for parsing the command line arguments and flags, and then
// executing the appropriate subcommand. It also sets up logging and error handling.
// The root command does not have any functionality of its own.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "spare",
		Short: "spare release single page application and aws infrastructure",
		Long: `,___,
[OvO] SPARE - Single Page Application Release Easily
/)__)         https://github.com/nao1215/spare (MIT LICENSE)
-"--"-`,
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newBugReportCmd())
	cmd.AddCommand(newInitCmd())
	cmd.AddCommand(newBuildCmd())
	return cmd
}

// Execute run process.
func Execute() int {
	if err := newRootCmd().Execute(); err != nil {
		log.Error(err)
		return 1
	}
	return 0
}
