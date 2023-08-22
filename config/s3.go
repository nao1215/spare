// Package config load settings from external files or environment variables and manage their values.
package config

import (
	"github.com/caarlos0/env/v9"
	"github.com/nao1215/spare/app/domain/model"
)

// S3 is a struct that contains the settings for the S3 storage.
type S3 struct {
	// Bucket is the name of the S3 bucket.
	Bucket model.Bucket `env:"spare_BUCKET"`
	// Region is the name of the AWS region.
	Region model.Region `env:"spare_REGION" envDefault:"us-east-1"`
}

// NewS3 returns a new S3 struct.
func NewS3() (*S3, error) {
	s3 := &S3{}
	if err := env.Parse(s3); err != nil {
		return nil, err
	}
	return s3, nil
}
