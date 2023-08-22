package config

import "errors"

var (
	// ErrInvalidRegion is an error that occurs when the region is invalid.
	ErrInvalidRegion = errors.New("invalid region")
	// ErrInvalidBucket is an error that occurs when the bucket is invalid.
	ErrInvalidBucket = errors.New("invalid bucket")
)
