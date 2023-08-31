// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/app/external"
	"github.com/nao1215/spare/app/interactor"
	"github.com/nao1215/spare/app/usecase"
)

// Injectors from wire.go:

// NewSpare returns a new Spare struct.
func NewSpare(profile model.AWSProfile, region model.Region, endpoint *model.Endpoint) (*Spare, error) {
	s3BucketCreator := external.NewS3BucketCreator(profile, region, endpoint)
	storageCreatorOptions := &interactor.StorageCreatorOptions{
		BucketCreator: s3BucketCreator,
	}
	storageCreator := interactor.NewStorageCreator(storageCreatorOptions)
	s3Uploader := external.NewS3Uploader(profile, region, endpoint)
	fileUploaderOptions := &interactor.FileUploaderOptions{
		FileUploader: s3Uploader,
	}
	fileUploader := interactor.NewFileUploader(fileUploaderOptions)
	spare := newSpare(storageCreator, fileUploader)
	return spare, nil
}

// wire.go:

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
