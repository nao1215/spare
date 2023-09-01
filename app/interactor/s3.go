// Package interactor is the implementation for usecase.
package interactor

import (
	"context"
	"errors"

	"github.com/charmbracelet/log"
	"github.com/google/wire"
	"github.com/nao1215/spare/app/domain/model"
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
	service.BucketPublicAccessBlocker
	service.BucketPolicySetter
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

	if _, err := s.opts.BucketPublicAccessBlocker.BlockBucketPublicAccess(ctx, &service.BucketPublicAccessBlockerInput{
		Bucket: input.BucketName,
		Region: input.Region,
	}); err != nil {
		return nil, err
	}

	if _, err := s.opts.BucketPolicySetter.SetBucketPolicy(ctx, &service.BucketPolicySetterInput{
		Bucket: input.BucketName,
		Policy: model.NewAllowCloudFrontS3BucketPolicy(input.BucketName),
	}); err != nil {
		return nil, err
	}

	return &usecase.CreateStorageOutput{}, nil
}

// FileUploaderSet is a provider set for FileUploader.
//
//nolint:gochecknoglobals
var FileUploaderSet = wire.NewSet(
	NewFileUploader,
	wire.Struct(new(FileUploaderOptions), "*"),
	wire.Bind(new(usecase.FileUploader), new(*FileUploader)),
)

var _ usecase.FileUploader = (*FileUploader)(nil)

// FileUploader is an implementation for FileUploader.
type FileUploader struct {
	opts *FileUploaderOptions
}

// FileUploaderOptions is an option struct for FileUploader.
type FileUploaderOptions struct {
	service.FileUploader
}

// NewFileUploader returns a new FileUploader struct.
func NewFileUploader(opts *FileUploaderOptions) *FileUploader {
	return &FileUploader{
		opts: opts,
	}
}

// UploadFile uploads a file to external storage.
func (u *FileUploader) UploadFile(ctx context.Context, input *usecase.UploadFileInput) (*usecase.UploadFileOutput, error) {
	if _, err := u.opts.FileUploader.UploadFile(ctx, &service.FileUploaderInput{
		BucketName: input.BucketName,
		Key:        input.Key,
		Data:       input.Data,
	}); err != nil {
		return nil, err
	}
	return &usecase.UploadFileOutput{}, nil
}
