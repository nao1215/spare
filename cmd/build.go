package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/charmbracelet/log"
	"github.com/nao1215/spare/app/di"
	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/app/usecase"
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
	cmd.Flags().StringP("profile", "p", "", "AWS profile name. if this is empty, use $AWS_PROFILE")
	return cmd
}

type builder struct {
	// ctx is a context.Context.
	ctx context.Context
	// spare is a struct that executes the build command.
	spare *di.Spare
	// config is a struct that contains the settings for the spare CLI command.
	config *config.Config
	// debug is a flag that indicates whether to run debug mode.
	debug bool
	// awsProfile is a profile name of AWS. If this is empty, use $AWS_PROFILE.
	awsProfile model.AWSProfile
}

// Parse parses the arguments and flags.
func (b *builder) Parse(cmd *cobra.Command, _ []string) (err error) {
	b.ctx = context.Background()
	if b.debug, err = cmd.Flags().GetBool("debug"); err != nil {
		return errfmt.Wrap(err, "can not parse command line argument (--debug)")
	}

	profile, err := cmd.Flags().GetString("profile")
	if err != nil {
		return errfmt.Wrap(err, "can not parse command line argument (--profile)")
	}
	b.awsProfile = model.NewAWSProfile(profile)

	b.config, err = b.readConfig()
	if err != nil {
		return err
	}
	var endpoint *model.Endpoint
	if b.debug {
		endpoint = &b.config.DebugLocalstackEndpoint
	}

	// Create a new instance of the Spare struct using the di.NewSpare function
	b.spare, err = di.NewSpare(b.awsProfile, b.config.Region, endpoint)
	if err != nil {
		return err
	}

	return nil
}

// Do generate .spare.yml at current directory.
// If .spare.yml already exists, return error.
func (b *builder) Do() error {
	log.Info("validate setting fron .spare.yml")
	if err := b.config.Validate(b.debug); err != nil {
		return err
	}
	log.Info("setting is valid")

	if err := b.confirm(); err != nil {
		return err
	}

	log.Info("start building AWS infrastructure")
	log.Info("create S3 bucket")
	if _, err := b.spare.StorageCreator.CreateStorage(b.ctx, &usecase.CreateStorageInput{
		BucketName: b.config.S3BucketName,
		Region:     b.config.Region,
	}); err != nil {
		return err
	}
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

// confirm shows the settings and asks if you want to build AWS infrastructure.
func (b *builder) confirm() error {
	fmt.Println("")
	fmt.Println("[DEBUG MODE]")
	fmt.Println(b.debug)
	fmt.Println("[AWS PROFILE]")
	fmt.Println(b.awsProfile)
	fmt.Println("[.spare.yml]")
	if err := b.config.Write(os.Stdout); err != nil {
		return err
	}
	fmt.Println("")

	var result bool
	if err := survey.AskOne(
		&survey.Confirm{
			Message: "want to build AWS infrastructure with the above settings?",
		},
		&result,
	); err != nil {
		return err
	}

	if !result {
		return errors.New("canceled")
	}
	return nil
}
