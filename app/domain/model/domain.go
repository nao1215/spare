package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nao1215/spare/errfmt"
)

// Domain is a type that represents a domain name.
type Domain string

// String returns the string representation of Domain.
func (d Domain) String() string {
	return string(d)
}

// Validate validates Domain. If Domain is invalid, it returns an error.
func (d Domain) Validate() error {
	if d == "" {
		return errfmt.Wrap(ErrInvalidDomain, "domain is empty")
	}
	for _, part := range strings.Split(d.String(), ".") {
		if len(part) == 0 || !isAlphaNumeric(part) {
			return errfmt.Wrap(ErrInvalidDomain, fmt.Sprintf("domain %s is invalid", d))
		}
	}
	return nil
}

// isAlphaNumeric　returns true if s is alphanumeric.
func isAlphaNumeric(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

// AllowOrigins is list of origins (domain names) that CloudFront can use as
// the value for the Access-Control-Allow-Origin HTTP response header.
type AllowOrigins []Domain

// Validate validates AllowOrigins. If AllowOrigins is invalid, it returns an error.
func (a AllowOrigins) Validate() (err error) {
	for _, origin := range a {
		if e := origin.Validate(); e != nil {
			err = errors.Join(err, e)
		}
	}
	return err
}

// Endpoint is a type that represents an endpoint.
type Endpoint string

// String returns the string representation of Endpoint.
func (e Endpoint) String() string {
	return string(e)
}

// DebugLocalstackEndpoint is the endpoint for localstack. It's used for testing.
const DebugLocalstackEndpoint = "http://localhost:4566"
