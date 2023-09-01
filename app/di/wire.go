//go:build wireinject
// +build wireinject

// Package di Inject dependence by wire command.
package di

import (
	"github.com/google/wire"
	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/app/external"
	"github.com/nao1215/spare/app/interactor"
	"github.com/nao1215/spare/app/usecase"
)

// NewSpare returns a new Spare struct.
func NewSpare(profile model.AWSProfile, region model.Region, endpoint *model.Endpoint) (*Spare, error) {
	wire.Build(
		interactor.StorageCreatorSet,
		interactor.FileUploaderSet,
		external.BuckerCreatorSet,
		external.FileUploaderSet,
		external.BucketPublicAccessBlockerSet,
		external.BucketPolicySetterSet,
		newSpare,
	)
	return nil, nil
}

// Spare is a struct that contains the settings for the spare CLI command.
type Spare struct {
	// StorageCreator is an interface for creating external storage.
	StorageCreator usecase.StorageCreator
	// FileUploader is an interface for uploading files to external storage.
	FileUploader usecase.FileUploader
}

// newSpare returns a new Spare struct.
func newSpare(
	storageCreator usecase.StorageCreator,
	fileUploader usecase.FileUploader,
) *Spare {
	return &Spare{
		StorageCreator: storageCreator,
		FileUploader:   fileUploader,
	}
}
