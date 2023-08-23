package config

import (
	"io"

	"github.com/nao1215/spare/app/domain/model"
	"gopkg.in/yaml.v2"
)

// Config is a struct that corresponds to the configuration file ".spare.yml".
type Config struct {
	// DeployTarget is the path of the deploy target (it's SPA).
	DeployTarget string `yaml:"deploy_target"`
	// Region is AWS region.
	Region model.Region `yaml:"region"`
	// CustomDomain is the domain name of the CloudFront.
	// If you do not specify this, the CloudFront default domain name is used.
	CustomDomain model.Domain `yaml:"custom_domain"`
	// S3BucketName is the name of the S3 bucket.
	S3BucketName string `yaml:"s3_bucket_name"`
	// CORS is the list of CORS configuration.
	CORS []model.Domain `yaml:"cors"`
	// TODO: WAF, HTTPS, Cache
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{
		DeployTarget: "src",
		Region:       model.RegionUSEast1,
		CustomDomain: "",
		S3BucketName: "",
		CORS:         []model.Domain{},
	}
}

// Write writes the Config to the io.Writer.
func (c *Config) Write(w io.Writer) error {
	encoder := yaml.NewEncoder(w)
	defer encoder.Close()
	return encoder.Encode(c)
}
