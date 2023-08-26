package config

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nao1215/spare/app/domain/model"
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

func TestConfig_Read(t *testing.T) {
	t.Parallel()

	t.Run("success to read yml data", func(t *testing.T) {
		t.Parallel()

		file, err := os.Open(filepath.Join("testdata", "read_test.yml"))
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			if closeErr := file.Close(); closeErr != nil {
				t.Fatal(closeErr)
			}
		}()

		got := NewConfig()
		if err := got.Read(file); err != nil {
			t.Fatal(err)
		}

		want := &Config{
			SpareTemplateVersion: "1.0.0",
			DeployTarget:         "test-src",
			Region:               model.RegionUSEast2,
			CustomDomain:         "example.com",
			S3BucketName:         "test-bucket",
			CORS:                 []model.Domain{"example.com", "test.example.com"},
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}
