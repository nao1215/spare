package config

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConfig_Write(t *testing.T) {
	t.Parallel()

	t.Run("success to write yml data", func(t *testing.T) {
		t.Parallel()

		c := NewConfig()
		want, err := os.ReadFile(filepath.Join("testdata", "test.yml"))
		if err != nil {
			t.Fatal(err)
		}

		got := bytes.NewBufferString("")
		if err := c.Write(got); err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got.String(), string(want)); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})

}
