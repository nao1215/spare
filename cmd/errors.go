package cmd

import (
	"fmt"

	"github.com/nao1215/spare/config"
)

var (
	// ErrConfigFileAlreadyExists is an error that occurs when the config file already exists.
	ErrConfigFileAlreadyExists = fmt.Errorf("%s config file already exists", config.ConfigFilePath)
)
