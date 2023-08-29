// Package interactor is the implementation for usecase.
package interactor

import (
	"context"

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
func (s *StorageCreator) CreateStorage(_ context.Context, _ *usecase.CreateStorageInput) (*usecase.CreateStorageOutput, error) {
	log.Info("not implemented yet")
	return &usecase.CreateStorageOutput{}, nil
}
