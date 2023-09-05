package cmd

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/nao1215/spare/app/di"
	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/config"
	"github.com/nao1215/spare/utils/errfmt"
	"github.com/spf13/cobra"
)

type commonOption struct {
	// ctx is a context.Context.
	ctx context.Context
	// spare is a struct that executes the sub command.
	spare *di.Spare
	// config is a struct that contains the settings for the spare CLI command.
	config *config.Config
	// debug is a flag that indicates whether to run debug mode.
	debug bool
	// configFilePath is a path of the config file.
	configFilePath string
	// awsProfile is a profile name of AWS. If this is empty, use $AWS_PROFILE.
	awsProfile model.AWSProfile
}

// Parse parses the arguments and flags.
func parseCommon(cmd *cobra.Command, _ []string) (*commonOption, error) {
	ctx := context.Background()
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return nil, errfmt.Wrap(err, "can not parse command line argument (--debug)")
	}

	configFilePath, err := cmd.Flags().GetString("file")
	if err != nil {
		return nil, errfmt.Wrap(err, "can not parse command line argument (--file)")
	}

	profile, err := cmd.Flags().GetString("profile")
	if err != nil {
		return nil, errfmt.Wrap(err, "can not parse command line argument (--profile)")
	}
	awsProfile := model.NewAWSProfile(profile)

	config, err := readConfig(configFilePath)
	if err != nil {
		return nil, err
	}

	var endpoint *model.Endpoint
	if debug {
		endpoint = &config.DebugLocalstackEndpoint
	}

	// Create a new instance of the Spare struct using the di.NewSpare function
	spare, err := di.NewSpare(awsProfile, config.Region, endpoint)
	if err != nil {
		return nil, err
	}
	return &commonOption{
		ctx:            ctx,
		spare:          spare,
		config:         config,
		configFilePath: configFilePath,
		debug:          debug,
		awsProfile:     awsProfile,
	}, nil
}

// readConfig reads .spare.yml and returns config.Config.
func readConfig(configFilePath string) (*config.Config, error) {
	file, err := os.Open(filepath.Clean(configFilePath))
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
