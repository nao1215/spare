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
		interactor.CDNCreatorSet,
		external.BuckerCreatorSet,
		external.FileUploaderSet,
		external.BucketPublicAccessBlockerSet,
		external.BucketPolicySetterSet,
		external.CDNCreatorSet,
		newSpare,
	)
	return nil, nil
}

// Spare is a struct that contains the settings for the spare CLI command.
type Spare struct {
	// StorageCreator is an interface for creating external storage.
	StorageCreator usecase.StorageCreator
	// CDNCreator is an interface for creating CDN.
	CDNCreator usecase.CDNCreator
	// FileUploader is an interface for uploading files to external storage.
	FileUploader usecase.FileUploader
}

// newSpare returns a new Spare struct.
func newSpare(
	storageCreator usecase.StorageCreator,
	cdncreator usecase.CDNCreator,
	fileUploader usecase.FileUploader,
) *Spare {
	return &Spare{
		StorageCreator: storageCreator,
		CDNCreator:     cdncreator,
		FileUploader:   fileUploader,
	}
}
