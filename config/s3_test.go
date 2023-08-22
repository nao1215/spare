// Package config load settings from external files or environment variables and manage their values.
package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewS3(t *testing.T) {
	tests := []struct {
		name    string
		want    *S3
		wantErr bool
	}{
		{
			name: "initialize success",
			want: &S3{
				Bucket: "spare",
				Region: "us-east-1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("spare_BUCKET", "spare")

			got, err := NewS3()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewS3() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("User value is mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
