package service

import (
	"context"

	"github.com/nao1215/spare/app/domain/model"
)

// CDNCreatorInput is an input struct for CDNCreator.
type CDNCreatorInput struct {
	// BucketName is the name of the  bucket.
	BucketName model.BucketName
}

// CDNCreatorOutput is an output struct for CDNCreator.
type CDNCreatorOutput struct {
	// Domain is the domain of the CDN.
	Domain model.Domain
}

// CDNCreator is an interface for creating CDN.
type CDNCreator interface {
	CreateCDN(context.Context, *CDNCreatorInput) (*CDNCreatorOutput, error)
}
