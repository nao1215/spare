package interactor

import (
	"context"

	"github.com/google/wire"
	"github.com/nao1215/spare/app/domain/service"
	"github.com/nao1215/spare/app/usecase"
)

// CDNCreatorSet is a set of CDNCreator.
//
//nolint:gochecknoglobals
var CDNCreatorSet = wire.NewSet(
	NewCDNCreator,
	wire.Struct(new(CDNCreatorOptions), "*"),
	wire.Bind(new(usecase.CDNCreator), new(*CDNCreator)),
)

var _ usecase.CDNCreator = (*CDNCreator)(nil)

// CDNCreator is an implementation for CDNCreator.
type CDNCreator struct {
	opts *CDNCreatorOptions
}

// CDNCreatorOptions is an option struct for CDNCreator.
type CDNCreatorOptions struct {
	service.CDNCreator
	service.OAICreator
}

// NewCDNCreator returns a new CDNCreator struct.
func NewCDNCreator(opts *CDNCreatorOptions) *CDNCreator {
	return &CDNCreator{
		opts: opts,
	}
}

// CreateCDN creates a CDN.
func (c *CDNCreator) CreateCDN(ctx context.Context, input *usecase.CreateCDNInput) (*usecase.CreateCDNOutput, error) {
	oaiOutput, err := c.opts.OAICreator.CreateOAI(ctx, &service.OAICreatorInput{})
	if err != nil {
		return nil, err
	}

	createCDNOutput, err := c.opts.CDNCreator.CreateCDN(ctx, &service.CDNCreatorInput{
		BucketName: input.BucketName,
		OAIID:      oaiOutput.ID,
	})
	if err != nil {
		return nil, err
	}
	return &usecase.CreateCDNOutput{
		Domain: createCDNOutput.Domain,
	}, nil
}
