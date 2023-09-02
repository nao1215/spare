// Package model contains the definitions of domain models and business logic.
package model

import (
	"errors"
	"strings"
	"testing"
)

func TestRegionString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		r    Region
		want string
	}{
		{
			name: "success",
			r:    RegionAPNortheast1,
			want: "ap-northeast-1",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.r.String(); got != tt.want {
				t.Errorf("Region.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegionValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		r       Region
		wantErr bool
		e       error
	}{
		{
			name:    "success",
			r:       RegionAPNortheast1,
			wantErr: false,
			e:       nil,
		},
		{
			name:    "failure. region is empty",
			r:       Region(""),
			wantErr: true,
			e:       ErrEmptyRegion,
		},
		{
			name:    "failure. region is invalid",
			r:       Region("invalid"),
			wantErr: true,
			e:       ErrInvalidRegion,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Region.Validate() error = %v, wantErr %v", err, tt.wantErr)
				if tt.wantErr {
					if errors.Is(err, tt.e) {
						t.Errorf("error mismatch got = %v, wantErr %v", err, tt.wantErr)
					}
				}
			}
		})
	}
}

func TestBucketString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		b    BucketName
		want string
	}{
		{
			name: "success",
			b:    BucketName("spare"),
			want: "spare",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.b.String(); got != tt.want {
				t.Errorf("Bucket.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBucketNameValidateLength(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		b       BucketName
		wantErr bool
	}{
		{
			name:    "success: minimum length",
			b:       BucketName("abc"),
			wantErr: false,
		},
		{
			name:    "success: maximum length",
			b:       BucketName(strings.Repeat("a", 63)),
			wantErr: false,
		},
		{
			name:    "failure. bucket name is too short",
			b:       BucketName("ab"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name is too long",
			b:       BucketName(strings.Repeat("a", 64)),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.b.validateLength(); (err != nil) != tt.wantErr {
				t.Errorf("BucketName.validateLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBucketNameValidatePattern(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		b       BucketName
		wantErr bool
	}{
		{
			name:    "success",
			b:       BucketName("abc"),
			wantErr: false,
		},
		{
			name:    "failure. bucket name contains invalid character",
			b:       BucketName("abc!"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name contains uppercase character",
			b:       BucketName("Abc"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name contains underscore",
			b:       BucketName("abc_def"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name starts with hyphen",
			b:       BucketName("-abc"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name ends with hyphen",
			b:       BucketName("abc-"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.b.validatePattern(); (err != nil) != tt.wantErr {
				t.Errorf("BucketName.validatePattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBucketNameValidatePrefix(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		b       BucketName
		wantErr bool
	}{
		{
			name:    "success",
			b:       BucketName("abc"),
			wantErr: false,
		},
		{
			name:    "failure. bucket name starts with 'xn--'",
			b:       BucketName("xn--abc"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name starts with 'sthree-'",
			b:       BucketName("sthree-abc"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name starts with 'sthree-configurator'",
			b:       BucketName("sthree-configurator-abc"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.b.validatePrefix(); (err != nil) != tt.wantErr {
				t.Errorf("BucketName.validatePrefix() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBucketNameValidateSuffix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		b       BucketName
		wantErr bool
	}{
		{
			name:    "success",
			b:       BucketName("abc"),
			wantErr: false,
		},
		{
			name:    "failure. bucket name ends with '-s3alias'",
			b:       BucketName("abc-s3alias"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name ends with '--ol-s3'",
			b:       BucketName("abc--ol-s3"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.b.validateSuffix(); (err != nil) != tt.wantErr {
				t.Errorf("BucketName.validateSuffix() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBucketNameValidateCharSequence(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		b       BucketName
		wantErr bool
	}{
		{
			name:    "success",
			b:       BucketName("abc"),
			wantErr: false,
		},
		{
			name:    "failure. bucket name contains consecutive periods",
			b:       BucketName("abc..def"),
			wantErr: true,
		},
		{
			name:    "failure. bucket name contains consecutive hyphens",
			b:       BucketName("abc--def"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.b.validateCharSequence(); (err != nil) != tt.wantErr {
				t.Errorf("BucketName.validateCharSequence() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBucketNameValidate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		b       BucketName
		wantErr bool
	}{
		{
			name:    "success",
			b:       BucketName("abc"),
			wantErr: false,
		},
		{
			name:    "failure. bucket name is empty",
			b:       BucketName(""),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.b.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("BucketName.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBucketNameDomain(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		b    BucketName
		want string
	}{
		{
			name: "success",
			b:    BucketName("abc"),
			want: "abc.s3.amazonaws.com",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.b.Domain(); got != tt.want {
				t.Errorf("BucketName.Domain() = %v, want %v", got, tt.want)
			}
		})
	}
}
