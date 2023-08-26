package config

import "testing"

func TestTemplateVersionString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		tr   TemplateVersion
		want string
	}{
		{
			name: "0.0.1",
			tr:   "0.0.1",
			want: "0.0.1",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := tt.tr.String(); got != tt.want {
				t.Errorf("TemplateVersion.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
