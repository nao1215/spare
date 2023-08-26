package config

import (
	"testing"
)

func TestDeployTargetString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		d    DeployTarget
		want string
	}{
		{
			name: "src",
			d:    "src",
			want: "src",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.d.String(); got != tt.want {
				t.Errorf("DeployTarget.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeployTargetValidate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		d       DeployTarget
		wantErr bool
	}{
		{
			name:    "success",
			d:       "src",
			wantErr: false,
		},
		{
			name:    "failure. deploy target is empty",
			d:       "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("DeployTarget.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
