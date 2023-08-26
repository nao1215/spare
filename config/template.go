package config

import "github.com/nao1215/spare/errfmt"

// TemplateVersion is a type that represents a spare template version.
type TemplateVersion string

// CurrentSpareTemplateVersion is the version of the template.
const CurrentSpareTemplateVersion TemplateVersion = "0.0.1"

// String returns the string representation of TemplateVersion.
func (t TemplateVersion) String() string {
	return string(t)
}

// Validate validates TemplateVersion. If TemplateVersion is invalid, it returns an error.
// TemplateVersion is invalid if it is empty.
func (t TemplateVersion) Validate() error {
	if t == "" {
		return errfmt.Wrap(ErrInvalidSpareTemplateVersion, "SpareTemplateVersion is empty")
	}
	return nil
}
