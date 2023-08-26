package model

import "errors"

var (
	// ErrInvalidRegion is an error that occurs when the region is invalid.
	ErrInvalidRegion = errors.New("invalid region")
	// ErrEmptyRegion is an error that occurs when the region is empty.
	ErrEmptyRegion = errors.New("region is empty")
	// ErrEmptyBucketName is an error that occurs when the bucket name is empty.
	ErrEmptyBucketName = errors.New("bucket name is empty")
	// ErrInvalidDomain is an error that occurs when the domain is invalid.
	ErrInvalidDomain = errors.New("invalid domain")
)
