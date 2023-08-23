package model

import "testing"

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
