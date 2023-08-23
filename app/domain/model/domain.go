package model

// Domain is a type that represents a domain name.
type Domain string

// String returns the string representation of Domain.
func (d Domain) String() string {
	return string(d)
}
