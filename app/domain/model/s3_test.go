// Package model contains the definitions of domain models and business logic.
package model

import (
	"errors"
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
			r:    Region("ap-northeast-1"),
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
			r:       Region("ap-northeast-1"),
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

func TestBucketValid(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		b       BucketName
		wantErr error
	}{
		{
			name:    "success",
			b:       BucketName("spare"),
			wantErr: nil,
		},
		{
			name:    "failure. bucket name is empty",
			b:       BucketName(""),
			wantErr: ErrEmptyBucketName,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.b.Validate(); !errors.Is(err, tt.wantErr) {
				t.Errorf("Bucket.Validate() error = %v, wantErr %v", err, tt.wantErr)
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
