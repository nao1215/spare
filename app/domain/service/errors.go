package service

import "errors"

var (
	// ErrBucketAlreadyExistsOwnedByOther is an error that occurs when the bucket already exists and is owned by another account.
	ErrBucketAlreadyExistsOwnedByOther = errors.New("bucket already exists and is owned by another account")
	// ErrBucketAlreadyOwnedByYou is an error that occurs when the bucket already exists and is owned by you.
	ErrBucketAlreadyOwnedByYou = errors.New("bucket already exists and is owned by you")
	// ErrBucketPublicAccessBlock is an error that occurs when the bucket public access block setting fails.
	ErrBucketPublicAccessBlock = errors.New("failed to set public access block")
	// ErrBucketPolicySet is an error that occurs when the bucket policy setting fails.
	ErrBucketPolicySet = errors.New("failed to set bucket policy")
	// ErrCDNAlreadyExist is an error that occurs when the CDN already exists.
	ErrCDNAlreadyExists = errors.New("CDN already exists")
)
