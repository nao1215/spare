package config

import (
	"errors"
	"fmt"
	"io"

	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/rand"
	"github.com/nao1215/spare/version"
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
	// AllowOrigins is the list of domains that are allowed to access the SPA.
	AllowOrigins            model.AllowOrigins `yaml:"allowOrigins"`
	DebugLocalstackEndpoint model.Endpoint     `yaml:"debugLocalstackEndpoint"`
	// TODO: WAF, HTTPS, Cache
}

// NewConfig returns a new Config.
func NewConfig() *Config {
	cfg := &Config{
		SpareTemplateVersion:    CurrentSpareTemplateVersion,
		DeployTarget:            "src",
		Region:                  model.RegionUSEast1,
		CustomDomain:            "",
		S3BucketName:            "",
		AllowOrigins:            model.AllowOrigins{},
		DebugLocalstackEndpoint: model.DebugLocalstackEndpoint,
	}
	cfg.S3BucketName = cfg.DefaultS3BucketName()
	return cfg
}

// DefaultS3BucketName returns the default S3 bucket name.
func (c *Config) DefaultS3BucketName() model.BucketName {
	const randomStrLen = 15
	return model.BucketName(
		fmt.Sprintf("%s-%s-%s",
			version.Name, c.Region, rand.RandomLowerAlphanumericStr(randomStrLen)))
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
// If debugMode is true, it validates the DebugLocalstackEndpoint.
func (c *Config) Validate(debugMode bool) error {
	validators := []model.Validator{
		c.SpareTemplateVersion,
		c.DeployTarget,
		c.Region,
		c.CustomDomain,
		c.S3BucketName,
		c.AllowOrigins,
	}
	if debugMode {
		validators = append(validators, c.DebugLocalstackEndpoint)
	}

	for _, v := range validators {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
