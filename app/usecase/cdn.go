package usecase

import (
	"context"

	"github.com/nao1215/spare/app/domain/model"
)

// CDNCreator is an interface for creating CDN.
type CDNCreator interface {
	CreateCDN(ctx context.Context, input *CreateCDNInput) (*CreateCDNOutput, error)
}

// CreateCDNInput is an input struct for CDNCreator.
type CreateCDNInput struct {
	// BucketName is the name of the  bucket.
	BucketName model.BucketName
}

// CreateCDNOutput is an output struct for CDNCreator.
type CreateCDNOutput struct {
	// Domain is the domain of the CDN.
	Domain model.Domain
}
