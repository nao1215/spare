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
		external.BuckerCreatorSet,
		newSpare,
	)
	return nil, nil
}

// Spare is a struct that contains the settings for the spare CLI command.
type Spare struct {
	// StorageCreator is an interface for creating external storage.
	StorageCreator usecase.StorageCreator
}

// newSpare returns a new Spare struct.
func newSpare(storageCreator usecase.StorageCreator) *Spare {
	return &Spare{
		StorageCreator: storageCreator,
	}
}
