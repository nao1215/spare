package service

import "errors"

var (
	// ErrBucketAlreadyExistsOwnedByOther is an error that occurs when the bucket already exists and is owned by another account.
	ErrBucketAlreadyExistsOwnedByOther = errors.New("bucket already exists and is owned by another account")
	// ErrBucketAlreadyOwnedByYou is an error that occurs when the bucket already exists and is owned by you.
	ErrBucketAlreadyOwnedByYou = errors.New("bucket already exists and is owned by you")
)
