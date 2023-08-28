package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/charmbracelet/log"
	"github.com/nao1215/spare/config"
	"github.com/nao1215/spare/utils/errfmt"
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
	log.Info("spare", "debug mode", b.debug)
	log.Info("validate setting fron .spare.yml")
	cfg, err := b.readConfig()
	if err != nil {
		return err
	}

	if err := cfg.Validate(b.debug); err != nil {
		return err
	}
	log.Info("setting is valid")

	if err := b.confirm(cfg); err != nil {
		return err
	}

	log.Info("start building AWS infrastructure")
	log.Info("create S3 bucket")
	// TODO: create S3 bucket and unit test

	return nil
}

// readConfig reads .spare.yml and returns config.Config.
func (b *builder) readConfig() (*config.Config, error) {
	file, err := os.Open(config.ConfigFilePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()

	cfg := config.NewConfig()
	if err := cfg.Read(file); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (b *builder) confirm(cfg *config.Config) error {
	fmt.Println("== .spare.yml ===================================")
	if err := cfg.Write(os.Stdout); err != nil {
		return err
	}
	fmt.Println("=================================================")

	var result bool
	if err := survey.AskOne(
		&survey.Confirm{
			Message: "want to build AWS infrastructure with the above settings?",
		},
		&result,
	); err != nil {
		return err
	}
	return nil
}
