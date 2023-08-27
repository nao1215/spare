package cmd

import (
	"github.com/nao1215/spare/errfmt"
	"github.com/spf13/cobra"
)

// newBuildCmd return build sub command.
func newBuildCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "build",
		Short:   "build AWS infrastructure for SPA",
		Example: "   spare build",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd, args, &builder{})
		},
	}
	cmd.Flags().BoolP("debug", "d", false, "run debug mode. you must run localstack before using this flag")
	return cmd
}

type builder struct {
	// debug is a flag that indicates whether to run debug mode.
	debug bool
}

// Parse parses the arguments and flags.
func (b *builder) Parse(cmd *cobra.Command, _ []string) (err error) {
	if b.debug, err = cmd.Flags().GetBool("debug"); err != nil {
		return errfmt.Wrap(err, "can not parse command line argument (--debug)")
	}
	return nil
}

// Do generate .spare.yml at current directory.
// If .spare.yml already exists, return error.
func (b *builder) Do() error {

	return nil
}
