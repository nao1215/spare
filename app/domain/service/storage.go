// Package service is an abstraction layer for accessing external services.
package service

import (
	"bytes"
	"context"
	"io"

	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/config"
)

// FileDownloderInput is an input struct for FileDownloader.
type FileDownloderInput struct {
	// Config is S3 Config.
	Config config.S3
	// Key is the S3 key
	Key string
}

// FileDownloderOutput is an output struct for FileDownloader.
type FileDownloderOutput struct {
	// Buffer is the downloaded data.
	*bytes.Buffer
}

// FileDownloder is an interface for download file to external storage.
type FileDownloder interface {
	// DownloadFile downloads a file from external storage.
	DownloadFile(context.Context, *FileDownloderInput) (*FileDownloderOutput, error)
}

// FileUploaderInput is an input struct for FileUploader.
type FileUploaderInput struct {
	// BucketName is the name of the bucket.
	BucketName model.BucketName
	// Key is the S3 key
	Key string
	// Data is the data to upload.
	Data io.Reader
}

// FileUploaderOutput is an output struct for FileUploader.
type FileUploaderOutput struct {
	// DetectedMIMEType is the MIME type detected by the library.
	DetectedMIMEType string
}

// FileUploader is an interface for uploading files to external storage.
type FileUploader interface {
	// UploadFile uploads a file from external storage.
	UploadFile(context.Context, *FileUploaderInput) (*FileUploaderOutput, error)
}

// BucketCreatorInput is an input struct for BucketCreator.
type BucketCreatorInput struct {
	// Bucket is the name of the  bucket.
	Bucket model.BucketName
	// Region is the name of the region.
	Region model.Region
}

// BucketCreatorOutput is an output struct for BucketCreator.
type BucketCreatorOutput struct{}

// BucketCreator is an interface for creating a bucket.
type BucketCreator interface {
	CreateBucket(context.Context, *BucketCreatorInput) (*BucketCreatorOutput, error)
}

// BucketPublicAccessBlockerInput is an input struct for BucketAccessBlocker.
type BucketPublicAccessBlockerInput struct {
	// Bucket is the name of the  bucket.
	Bucket model.BucketName
	// Region is the name of the region.
	Region model.Region
}

// BucketPublicAccessBlockerOutput is an output struct for BucketAccessBlocker.
type BucketPublicAccessBlockerOutput struct{}

// BucketPublicAccessBlocker is an interface for blocking access to a bucket.
type BucketPublicAccessBlocker interface {
	BlockBucketPublicAccess(context.Context, *BucketPublicAccessBlockerInput) (*BucketPublicAccessBlockerOutput, error)
}

// BucketPolicySetterInput is an input struct for BucketPolicySetter.
type BucketPolicySetterInput struct {
	// Bucket is the name of the  bucket.
	Bucket model.BucketName
	// Policy is the policy to set.
	Policy *model.BucketPolicy
}

// BucketPolicySetterOutput is an output struct for BucketPolicySetter.
type BucketPolicySetterOutput struct{}

// BucketPolicySetter is an interface for setting a bucket policy.
type BucketPolicySetter interface {
	SetBucketPolicy(context.Context, *BucketPolicySetterInput) (*BucketPolicySetterOutput, error)
}
