package model

import (
	"errors"
	"testing"
)

func TestDomainString(t *testing.T) {
	t.Parallel()

	const exampleCom = "example.com"
	tests := []struct {
		name string
		d    Domain
		want string
	}{
		{
			name: exampleCom,
			d:    exampleCom,
			want: exampleCom,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Domain.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDomainValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		d       Domain
		wantErr error
	}{
		{
			name:    "success",
			d:       "example.com",
			wantErr: nil,
		},
		{
			name:    "failure. protocol is included",
			d:       "https://example.com",
			wantErr: ErrInvalidDomain,
		},
		{
			name:    "failure. domain is empty",
			d:       "",
			wantErr: ErrInvalidDomain,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.d.Validate(); !errors.Is(err, tt.wantErr) {
				t.Errorf("Domain.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	t.Parallel()

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{s: "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"},
			want: true,
		},
		{
			name: "failure",
			args: args{s: "abc123/"},
			want: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isAlphaNumeric(tt.args.s); got != tt.want {
				t.Errorf("isAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllowOriginsValidate(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		a       AllowOrigins
		wantErr bool
	}{
		{
			name:    "success",
			a:       AllowOrigins{"example.com", "example.net"},
			wantErr: false,
		},
		{
			name:    "failure. origin is empty",
			a:       AllowOrigins{"example.com", ""},
			wantErr: true,
		},
		{
			name:    "failure. origin is invalid",
			a:       AllowOrigins{"example.com", "https://example.com"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.a.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("AllowOrigins.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEndpointString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		e    Endpoint
		want string
	}{
		{
			name: "success",
			e:    "http://example.com",
			want: "http://example.com",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.e.String(); got != tt.want {
				t.Errorf("Endpoint.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndpointValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		e       Endpoint
		wantErr bool
	}{
		{
			name:    "success",
			e:       "http://example.com",
			wantErr: false,
		},
		{
			name:    "failure. protocol is not included",
			e:       "example.com",
			wantErr: true,
		},
		{
			name:    "failure. endpoint is empty",
			e:       "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := tt.e.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Endpoint.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
