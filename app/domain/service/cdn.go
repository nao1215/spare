package service

import (
	"context"

	"github.com/nao1215/spare/app/domain/model"
)

// CDNCreatorInput is an input struct for CDNCreator.
type CDNCreatorInput struct {
	// BucketName is the name of the  bucket.
	BucketName model.BucketName
	// OAIID is the ID of the OAI.
	OAIID *string
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

// OAICreatorInput is an input struct for OAICreator.
type OAICreatorInput struct {
}

// OAICreatorOutput is an output struct for OAICreator.
type OAICreatorOutput struct {
	// ID is the ID of the OAI.
	ID *string
}

// OAICreator is an interface for creating OAI.
// OAI is an Origin Access Identity.
type OAICreator interface {
	CreateOAI(context.Context, *OAICreatorInput) (*OAICreatorOutput, error)
}
