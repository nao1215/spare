// Package interactor is the implementation for usecase.
package interactor

import (
	"context"
	"errors"

	"github.com/charmbracelet/log"
	"github.com/google/wire"
	"github.com/nao1215/spare/app/domain/service"
	"github.com/nao1215/spare/app/usecase"
)

// StorageCreatorSet is a provider set for StorageCreator.
//
//nolint:gochecknoglobals
var StorageCreatorSet = wire.NewSet(
	NewStorageCreator,
	wire.Struct(new(StorageCreatorOptions), "*"),
	wire.Bind(new(usecase.StorageCreator), new(*StorageCreator)),
)

var _ usecase.StorageCreator = (*StorageCreator)(nil)

// StorageCreator is an implementation for StorageCreator.
type StorageCreator struct {
	opts *StorageCreatorOptions
}

// StorageCreatorOptions is an option struct for StorageCreator.
type StorageCreatorOptions struct {
	service.BucketCreator
}

// NewStorageCreator returns a new StorageCreator struct.
func NewStorageCreator(opts *StorageCreatorOptions) *StorageCreator {
	return &StorageCreator{
		opts: opts,
	}
}

// CreateStorage creates a new external storage.
func (s *StorageCreator) CreateStorage(ctx context.Context, input *usecase.CreateStorageInput) (*usecase.CreateStorageOutput, error) {
	if _, err := s.opts.BucketCreator.CreateBucket(ctx, &service.BucketCreatorInput{
		Bucket: input.BucketName,
		Region: input.Region,
	}); err != nil {
		if errors.Is(err, service.ErrBucketAlreadyOwnedByYou) {
			// not error.
			log.Info("you already create the bucket", "bucket name", input.BucketName.String())
		} else {
			return nil, err
		}
	}
	return &usecase.CreateStorageOutput{}, nil
}
