package config

import "errors"

var (
	// ErrInvalidRegion is an error that occurs when the region is invalid.
	ErrInvalidRegion = errors.New("invalid region")
	// ErrInvalidBucket is an error that occurs when the bucket is invalid.
	ErrInvalidBucket = errors.New("invalid bucket")
	// ErrInvalidSpareTemplateVersion is an error that occurs when the spare template version is invalid.
	ErrInvalidSpareTemplateVersion = errors.New("invalid spare template version")
	// ErrInvalidDeployTarget is an error that occurs when the deploy target is invalid.
	ErrInvalidDeployTarget = errors.New("invalid deploy target")
)
