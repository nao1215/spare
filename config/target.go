package config

import "github.com/nao1215/spare/errfmt"

// DeployTarget is a type that represents a deploy target path.
type DeployTarget string

// String returns the string representation of DeployTarget.
func (d DeployTarget) String() string {
	return string(d)
}

// Validate validates DeployTarget. If DeployTarget is invalid, it returns an error.
// DeployTarget is invalid if it is empty.
func (d DeployTarget) Validate() error {
	if d == "" {
		return errfmt.Wrap(ErrInvalidDeployTarget, "DeployTarget is empty")
	}
	// TODO: check if the path exists
	return nil
}
