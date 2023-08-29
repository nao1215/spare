package model

// Validator is an interface that represents a validator.
type Validator interface {
	// Validate validates the value.
	Validate() error
}

// ValidationFunc is a type that represents a validation function.
type ValidationFunc func() error

// BucketNamer is an interface that represents a bucket name.
type BucketNamer interface {
	// String returns the string representation of BucketName.
	String() string
	// Empty is whether bucket name is empty
	Empty() bool
	// Validate validates BucketName. If BucketName is invalid, it returns an error.
	Validate() error
}

// Regioner is an interface that represents a region.
type Regioner interface {
	// String returns the string representation of Region.
	String() string
	// Validate validates Region. If Region is invalid, it returns an error.
	Validate() error
}
