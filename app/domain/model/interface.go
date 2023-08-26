package model

// Validator is an interface that represents a validator.
type Validator interface {
	// Validate validates the value.
	Validate() error
}
