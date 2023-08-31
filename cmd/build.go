package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/charmbracelet/log"
	"github.com/nao1215/spare/app/di"
	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/app/usecase"
	"github.com/nao1215/spare/config"
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
	commonOption, err := parseCommon(cmd, nil)
	if err != nil {
		return err
	}
	b.ctx = commonOption.ctx
	b.spare = commonOption.spare
	b.config = commonOption.config
	b.debug = commonOption.debug
	b.awsProfile = commonOption.awsProfile

	return nil
}

// Do generate .spare.yml at current directory.
// If .spare.yml already exists, return error.
func (b *builder) Do() error {
	log.Info("[VALIDATE] check .spare.yml")
	if err := b.config.Validate(b.debug); err != nil {
		return err
	}
	log.Info("[VALIDATE] ok .spare.yml")

	if err := b.confirm(); err != nil {
		return err
	}

	log.Info("[ CREATE ] start building AWS infrastructure")
	log.Info("[ CREATE ] s3 bucket", "name", b.config.S3BucketName.String())
	if _, err := b.spare.StorageCreator.CreateStorage(b.ctx, &usecase.CreateStorageInput{
		BucketName: b.config.S3BucketName,
		Region:     b.config.Region,
	}); err != nil {
		return err
	}

	return nil
}

// confirm shows the settings and asks if you want to build AWS infrastructure.
func (b *builder) confirm() error {
	log.Info("[CONFIRM ] check the settings")
	fmt.Println("")
	fmt.Println("[debug mode]")
	fmt.Printf(" %t\n", b.debug)
	fmt.Println("[aws profile]")
	fmt.Printf(" %s\n", b.awsProfile.String())
	fmt.Println("[.spare.yml]")
	fmt.Printf(" spareTemplateVersion: %s\n", b.config.SpareTemplateVersion)
	fmt.Printf(" deployTarget: %s\n", b.config.DeployTarget)
	fmt.Printf(" region: %s\n", b.config.Region)
	fmt.Printf(" customDomain: %s\n", b.config.CustomDomain)
	fmt.Printf(" s3BucketName: %s\n", b.config.S3BucketName)
	fmt.Printf(" allowOrigins: %s\n", b.config.AllowOrigins.String())
	if b.debug {
		fmt.Printf(" debugLocalstackEndpoint: %s\n", b.config.DebugLocalstackEndpoint)
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
