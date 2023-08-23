package cmd

import (
	"os"

	"github.com/nao1215/gorky/file"
	"github.com/nao1215/spare/config"
	"github.com/spf13/cobra"
)

// newInitCmd return init sub command.
func newInitCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "init",
		Short:   "Generate .spare.yml at current directory",
		Example: "   spare init",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd, args, &initializer{})
		},
	}
}

type initializer struct {
}

// Parse parses the arguments and flags.
func (i *initializer) Parse(_ *cobra.Command, _ []string) error {
	return nil
}

// Do generate .spare.yml at current directory.
// If .spare.yml already exists, return error.
func (i *initializer) Do(_ *cobra.Command) error {
	if file.IsFile(config.ConfigFilePath) {
		return ErrConfigFileAlreadyExists
	}

	file, err := os.Create(config.ConfigFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	config := config.NewConfig()
	return config.Write(file)
}
