package config

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConfigWrite(t *testing.T) {
	t.Parallel()

	t.Run("success to write yml data", func(t *testing.T) {
		t.Parallel()

		c := NewConfig()
		testFile := filepath.Join("testdata", "test.yml")
		if runtime.GOOS == "windows" {
			testFile = filepath.Join("testdata", "test_windows.yml")
		}

		want, err := os.ReadFile(filepath.Clean(testFile))
		if err != nil {
			t.Fatal(err)
		}

		got := bytes.NewBufferString("")
		if err := c.Write(got); err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(strings.ReplaceAll(got.String(), "\r", ""), strings.ReplaceAll(string(want), "\r", "")); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}
