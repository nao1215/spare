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

const (
	exampleCom                  = "example.com"
	exampleComWithTestSubDomain = "test.example.com"
	exampleComWithProtocol      = "https://example.com"
	testBucketName              = "test-bucket"
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

func TestConfigRead(t *testing.T) {
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
			SpareTemplateVersion:    "1.0.0",
			DeployTarget:            "test-src",
			Region:                  model.RegionUSEast2,
			CustomDomain:            exampleCom,
			S3BucketName:            testBucketName,
			AllowOrigins:            model.AllowOrigins{exampleCom, exampleComWithTestSubDomain},
			DebugLocalstackEndpoint: model.DebugLocalstackEndpoint,
		}

		if diff := cmp.Diff(want, got); diff != "" {
			t.Errorf("value is mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestConfigValidate(t *testing.T) {
	t.Parallel()
	type fields struct {
		SpareTemplateVersion TemplateVersion
		DeployTarget         DeployTarget
		Region               model.Region
		CustomDomain         model.Domain
		S3BucketName         model.BucketName
		AllowOrigins         model.AllowOrigins
		Endpoint             model.Endpoint
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				SpareTemplateVersion: "1.0.0",
				DeployTarget:         "src",
				Region:               model.RegionUSEast1,
				CustomDomain:         exampleCom,
				S3BucketName:         testBucketName,
				AllowOrigins:         model.AllowOrigins{exampleCom, exampleComWithTestSubDomain},
				Endpoint:             model.DebugLocalstackEndpoint,
			},
			wantErr: false,
		},
		{
			name: "failure. SpareTemplateVersion is empty",
			fields: fields{
				SpareTemplateVersion: "",
				DeployTarget:         "src",
				Region:               model.RegionUSEast1,
				CustomDomain:         exampleCom,
				S3BucketName:         testBucketName,
				AllowOrigins:         model.AllowOrigins{exampleCom, exampleComWithTestSubDomain},
				Endpoint:             model.DebugLocalstackEndpoint,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &Config{
				SpareTemplateVersion:    tt.fields.SpareTemplateVersion,
				DeployTarget:            tt.fields.DeployTarget,
				Region:                  tt.fields.Region,
				CustomDomain:            tt.fields.CustomDomain,
				S3BucketName:            tt.fields.S3BucketName,
				AllowOrigins:            tt.fields.AllowOrigins,
				DebugLocalstackEndpoint: tt.fields.Endpoint,
			}
			if err := c.Validate(false); (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
