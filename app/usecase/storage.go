// Package usecase is the abstraction layer for the business logic.
package usecase

import (
	"context"
	"io"

	"github.com/nao1215/spare/app/domain/model"
)

// StorageCreator is an interface for creating external storage.
type StorageCreator interface {
	CreateStorage(ctx context.Context, input *CreateStorageInput) (*CreateStorageOutput, error)
}

// CreateStorageInput is an input struct for StorageCreator.
type CreateStorageInput struct {
	// BucketName is the name of the bucket.
	BucketName model.BucketName
	// Region is the name of the region where the bucket is located.
	Region model.Region
}

// CreateStorageOutput is an output struct for StorageCreator.
type CreateStorageOutput struct{}

// FileUploader is an interface for uploading files to external storage.
type FileUploader interface {
	// UploadFile uploads a file from external storage.
	UploadFile(ctx context.Context, input *UploadFileInput) (*UploadFileOutput, error)
}

// UploadFileInput is an input struct for FileUploader.
type UploadFileInput struct {
	// BucketName is the name of the bucket.
	BucketName model.BucketName
	// Region is the name of the region where the bucket is located.
	Region model.Region
	// Key is the S3 key
	Key string
	// Data is the data to upload.
	Data io.Reader
}

// UploadFileOutput is an output struct for FileUploader.
type UploadFileOutput struct{}
