package model

import "os"

// AWSProfile is the name of the AWS profile.
type AWSProfile string

// NewAWSProfile returns a new AWSProfile.
// If p is empty, read $AWS_PROFILE and return it.
func NewAWSProfile(p string) AWSProfile {
	if p == "" {
		profile := os.Getenv("AWS_PROFILE")
		if profile == "" {
			return AWSProfile("default")
		}
		return AWSProfile(profile)
	}
	return AWSProfile(p)
}

// String returns the string representation of the AWSProfile.
func (p AWSProfile) String() string {
	return string(p)
}
