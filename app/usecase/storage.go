package usecase

import (
	"context"

	"github.com/nao1215/spare/app/domain/model"
)

// StorageCreator is an interface for creating external storage.
type StorageCreator interface {
	CreateStorage(ctx context.Context, input *CreateStorageInput) (*CreateStorageOutput, error)
}

// CreateStorageInput is an input struct for StorageCreator.
type CreateStorageInput struct {
	// BucketName is the name of the bucket.
	BucketName model.BucketNamer
	// Region is the name of the region where the bucket is located.
	Region model.Regioner
}

// CreateStorageOutput is an output struct for StorageCreator.
type CreateStorageOutput struct{}
