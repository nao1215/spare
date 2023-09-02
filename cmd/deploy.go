package cmd

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/nao1215/spare/app/di"
	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/app/usecase"
	"github.com/nao1215/spare/config"
	"github.com/nao1215/spare/utils/file"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

// newDeployCmd return deploy sub command.
func newDeployCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deploy",
		Short:   "deploy SPA to AWS",
		Example: "   spare deploy",
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(cmd, args, &deployer{})
		},
	}
	cmd.Flags().BoolP("debug", "d", false, "run debug mode. you must run localstack before using this flag")
	cmd.Flags().StringP("profile", "p", "", "AWS profile name. if this is empty, use $AWS_PROFILE")
	return cmd
}

type deployer struct {
	// ctx is a context.Context.
	ctx context.Context
	// spare is a struct that executes the deploy command.
	spare *di.Spare
	// config is a struct that contains the settings for the spare CLI command.
	config *config.Config
	// debug is a flag that indicates whether to run debug mode.
	debug bool
	// awsProfile is a profile name of AWS. If this is empty, use $AWS_PROFILE.
	awsProfile model.AWSProfile
}

// Parse parses the arguments and flags.
func (d *deployer) Parse(cmd *cobra.Command, _ []string) (err error) {
	commonOption, err := parseCommon(cmd, nil)
	if err != nil {
		return err
	}
	d.ctx = commonOption.ctx
	d.spare = commonOption.spare
	d.config = commonOption.config
	d.debug = commonOption.debug
	d.awsProfile = commonOption.awsProfile

	return nil
}

func (d *deployer) Do() error {
	log.Info("[  MODE  ]", "debug", d.debug)
	log.Info("[ CONFIG ]", "profile", d.awsProfile)
	log.Info("[ DEPLOY ]", "target path", d.config.DeployTarget, "bucket name", d.config.S3BucketName)

	files, err := file.WalkDir(d.config.DeployTarget.String())
	if err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(d.ctx)
	weighted := semaphore.NewWeighted(int64(runtime.NumCPU()))
	for _, file := range files {
		file := file
		eg.Go(func() error {
			if err := weighted.Acquire(ctx, 1); err != nil {
				return err
			}
			defer weighted.Release(1)

			f, err := os.Open(filepath.Clean(file))
			if err != nil {
				return err
			}
			defer func() {
				if closeErr := f.Close(); closeErr != nil {
					err = closeErr
				}
			}()

			key := strings.Replace(file, d.config.DeployTarget.String()+string(filepath.Separator), "", 1)
			if _, err := d.spare.FileUploader.UploadFile(ctx, &usecase.UploadFileInput{
				BucketName: d.config.S3BucketName,
				Region:     d.config.Region,
				// e.g. src/index.html -> index.html
				Key:  key,
				Data: f,
			}); err != nil {
				return err
			}
			log.Info("[ DEPLOY ]", "file name", key)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}
