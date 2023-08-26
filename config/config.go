package config

import (
	"errors"
	"io"

	"github.com/nao1215/spare/app/domain/model"
	"gopkg.in/yaml.v2"
)

// ConfigFilePath is the path of the configuration file.
const ConfigFilePath string = ".spare.yml"

// Config is a struct that corresponds to the configuration file ".spare.yml".
type Config struct {
	SpareTemplateVersion TemplateVersion `yaml:"spareTemplateVersion"`
	// DeployTarget is the path of the deploy target (it's SPA).
	DeployTarget DeployTarget `yaml:"deployTarget"`
	// Region is AWS region.
	Region model.Region `yaml:"region"`
	// CustomDomain is the domain name of the CloudFront.
	// If you do not specify this, the CloudFront default domain name is used.
	CustomDomain model.Domain `yaml:"customDomain"`
	// S3BucketName is the name of the S3 bucket.
	S3BucketName model.BucketName `yaml:"s3BucketName"`
	// CORS is the list of CORS configuration.
	CORS []model.Domain `yaml:"cors"`
	// TODO: WAF, HTTPS, Cache
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	return &Config{
		SpareTemplateVersion: CurrentSpareTemplateVersion,
		DeployTarget:         "src",
		Region:               model.RegionUSEast1,
		CustomDomain:         "",
		S3BucketName:         "",
		CORS:                 []model.Domain{},
	}
}

// Write writes the Config to the io.Writer.
func (c *Config) Write(w io.Writer) (err error) {
	encoder := yaml.NewEncoder(w)
	defer func() {
		if closeErr := encoder.Close(); closeErr != nil {
			err = errors.Join(err, closeErr)
		}
	}()
	return encoder.Encode(c)
}

// Read reads the Config from the io.Reader.
func (c *Config) Read(r io.Reader) (err error) {
	decoder := yaml.NewDecoder(r)
	return decoder.Decode(c)
}

// Validate validates the Config.
func (c *Config) Validate() error {
	validators := []Validator{
		c.SpareTemplateVersion,
		c.DeployTarget,
		c.Region,
		c.S3BucketName,
	}
	for _, v := range validators {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
