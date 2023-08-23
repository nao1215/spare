// Package cmd is a package that contains subcommands for the spare CLI command.
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// newRootCmd creates a new root command. This command is the entry point of the CLI.
// It is responsible for parsing the command line arguments and flags, and then
// executing the appropriate subcommand. It also sets up logging and error handling.
// The root command does not have any functionality of its own.
func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "spare",
		Short: "spare deploy single page application and aws infrastructure",
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.SilenceUsage = true
	cmd.SilenceErrors = true
	cmd.AddCommand(newVersionCmd())
	cmd.AddCommand(newBugReportCmd())
	cmd.AddCommand(newInitCmd())
	return cmd
}

// Execute run process.
func Execute() int {
	if err := newRootCmd().Execute(); err != nil {
		fmt.Printf("%s %s\n", color.RedString("[ERROR]"), err.Error())
		return 1
	}
	return 0
}
