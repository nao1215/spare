// Package external is the implementation for accessing external services.
package external

import (
	"bytes"
	"context"
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/wire"
	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/app/domain/service"
	"github.com/nao1215/spare/config"
	"github.com/nao1215/spare/utils/errfmt"
)

// S3Downloader is an implementation for FileDownloader.
type S3Downloader struct {
	*s3manager.Downloader
}

var _ service.FileDownloder = &S3Downloader{}

// downloadBufferSize is the buffer size for downloading files. It's 5MB.
const downloadBufferSize = 5 * 1024 * 1024

// NewS3Downloader returns a new S3Downloader struct.
func NewS3Downloader(config config.S3) *S3Downloader {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, // Ref. ~/.aws/config
		Config:            aws.Config{Region: aws.String(config.Region.String())},
	}))
	downloader := s3manager.NewDownloader(session, func(d *s3manager.Downloader) {
		d.BufferProvider = s3manager.NewPooledBufferedWriterReadFromProvider(downloadBufferSize)
	})
	return &S3Downloader{downloader}
}

// DownloadFile downloads a file from S3.
func (s *S3Downloader) DownloadFile(_ context.Context, input *service.FileDownloderInput) (*service.FileDownloderOutput, error) {
	buf := aws.NewWriteAtBuffer([]byte{})
	objInput := &s3.GetObjectInput{
		Bucket: aws.String(input.Config.Bucket.String()),
		Key:    aws.String(input.Key),
	}

	if _, err := s.Download(buf, objInput); err != nil {
		return nil, err
	}
	return &service.FileDownloderOutput{
		Buffer: bytes.NewBuffer(buf.Bytes()),
	}, nil
}

// FileUploaderSet is a provider set for FileUploader.
//
//nolint:gochecknoglobals
var FileUploaderSet = wire.NewSet(
	NewS3Uploader,
	wire.Bind(new(service.FileUploader), new(*S3Uploader)),
)

// S3Uploader is an implementation for FileUploader.
type S3Uploader struct {
	*s3manager.Uploader
}

var _ service.FileUploader = &S3Uploader{}

// NewS3Uploader returns a new S3Uploader struct.
func NewS3Uploader(profile model.AWSProfile, region model.Region, endpoint *model.Endpoint) *S3Uploader {
	return &S3Uploader{s3manager.NewUploader(newS3Session(profile, region, endpoint))}
}

// UploadFile uploads a file to S3.
func (s *S3Uploader) UploadFile(_ context.Context, input *service.FileUploaderInput) (*service.FileUploaderOutput, error) {
	contentDetectReader, uploadReader, err := duplicateReader(input.Data)
	if err != nil {
		return nil, errfmt.Wrap(service.ErrFileUpload, err.Error())
	}
	contentType, err := detectContentType(contentDetectReader, input.Key)
	if err != nil {
		return nil, errfmt.Wrap(service.ErrFileUpload, err.Error())
	}

	uploadInput := &s3manager.UploadInput{
		Bucket:      aws.String(input.BucketName.String()),
		Body:        aws.ReadSeekCloser(uploadReader),
		Key:         aws.String(input.Key),
		ContentType: aws.String(contentType),
	}

	if _, err := s.Upload(uploadInput); err != nil {
		return nil, err
	}
	return &service.FileUploaderOutput{
		DetectedMIMEType: contentType,
	}, nil
}

// BuckerCreatorSet is a provider set for BuckerCreator.
//
//nolint:gochecknoglobals
var BuckerCreatorSet = wire.NewSet(
	NewS3BucketCreator,
	wire.Bind(new(service.BucketCreator), new(*S3BucketCreator)),
)

// S3BucketCreator is an implementation for BucketCreator.
type S3BucketCreator struct {
	svc *s3.S3
}

var _ service.BucketCreator = &S3BucketCreator{}

// NewS3BucketCreator returns a new S3BucketCreator struct.
func NewS3BucketCreator(profile model.AWSProfile, region model.Region, endpoint *model.Endpoint) *S3BucketCreator {
	return &S3BucketCreator{s3.New(newS3Session(profile, region, endpoint))}
}

// CreateBucket creates a bucket on S3.
func (s *S3BucketCreator) CreateBucket(_ context.Context, input *service.BucketCreatorInput) (*service.BucketCreatorOutput, error) {
	createBucketConfig := &s3.CreateBucketConfiguration{}
	createBucketConfig.SetLocationConstraint(input.Region.String())

	if _, err := s.svc.CreateBucket(&s3.CreateBucketInput{
		Bucket:                    aws.String(input.Bucket.String()),
		CreateBucketConfiguration: createBucketConfig,
	}); err != nil {
		var awsErr awserr.Error
		if errors.As(err, &awsErr) {
			switch awsErr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				return nil, service.ErrBucketAlreadyExistsOwnedByOther
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				return nil, service.ErrBucketAlreadyOwnedByYou
			default:
				return nil, err
			}
		}
		return nil, err
	}
	return &service.BucketCreatorOutput{}, nil
}

// BucketPublicAccessBlockerSet is a provider set for BucketPublicAccessBlocker.
//
//nolint:gochecknoglobals
var BucketPublicAccessBlockerSet = wire.NewSet(
	NewS3BucketPublicAccessBlocker,
	wire.Bind(new(service.BucketPublicAccessBlocker), new(*S3BucketPublicAccessBlocker)),
)

// S3BucketPublicAccessBlocker is an implementation for BucketPublicAccessBlocker.
type S3BucketPublicAccessBlocker struct {
	svc *s3.S3
}

var _ service.BucketPublicAccessBlocker = &S3BucketPublicAccessBlocker{}

// NewS3BucketPublicAccessBlocker returns a new S3BucketPublicAccessBlocker struct.
func NewS3BucketPublicAccessBlocker(profile model.AWSProfile, region model.Region, endpoint *model.Endpoint) *S3BucketPublicAccessBlocker {
	return &S3BucketPublicAccessBlocker{s3.New(newS3Session(profile, region, endpoint))}
}

// BlockBucketPublicAccess blocks public access to a bucket on S3.
func (s *S3BucketPublicAccessBlocker) BlockBucketPublicAccess(_ context.Context, input *service.BucketPublicAccessBlockerInput) (*service.BucketPublicAccessBlockerOutput, error) {
	_, err := s.svc.PutPublicAccessBlock(&s3.PutPublicAccessBlockInput{
		Bucket: aws.String(input.Bucket.String()),
		PublicAccessBlockConfiguration: &s3.PublicAccessBlockConfiguration{
			BlockPublicAcls:       aws.Bool(true),
			BlockPublicPolicy:     aws.Bool(true),
			IgnorePublicAcls:      aws.Bool(true),
			RestrictPublicBuckets: aws.Bool(true),
		},
	})
	if err != nil {
		return nil, errfmt.Wrap(service.ErrBucketPublicAccessBlock, err.Error())
	}
	return &service.BucketPublicAccessBlockerOutput{}, nil
}

// BucketPolicySetterSet is a provider set for BucketPolicySetter.
//
//nolint:gochecknoglobals
var BucketPolicySetterSet = wire.NewSet(
	NewS3BucketPolicySetter,
	wire.Bind(new(service.BucketPolicySetter), new(*S3BucketPolicySetter)),
)

// S3BucketPolicySetter is an implementation for BucketPolicySetter.
type S3BucketPolicySetter struct {
	svc *s3.S3
}

var _ service.BucketPolicySetter = &S3BucketPolicySetter{}

// NewS3BucketPolicySetter returns a new S3BucketPolicySetter struct.
func NewS3BucketPolicySetter(profile model.AWSProfile, region model.Region, endpoint *model.Endpoint) *S3BucketPolicySetter {
	return &S3BucketPolicySetter{s3.New(newS3Session(profile, region, endpoint))}
}

// SetBucketPolicy sets a bucket policy on S3.
func (s *S3BucketPolicySetter) SetBucketPolicy(_ context.Context, input *service.BucketPolicySetterInput) (*service.BucketPolicySetterOutput, error) {
	policy, err := input.Policy.String()
	if err != nil {
		return nil, err
	}
	_, err = s.svc.PutBucketPolicy(&s3.PutBucketPolicyInput{
		Bucket: aws.String(input.Bucket.String()),
		Policy: aws.String(policy),
	})
	if err != nil {
		return nil, errfmt.Wrap(service.ErrBucketPolicySet, err.Error())
	}
	return &service.BucketPolicySetterOutput{}, nil
}
