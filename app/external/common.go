package external

import (
	"bytes"
	"io"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gabriel-vasile/mimetype"
	"github.com/nao1215/spare/app/domain/model"
	"github.com/nao1215/spare/app/domain/service"
	"github.com/nao1215/spare/utils/errfmt"
)

// newS3Session returns a new session.
func newS3Session(profile model.AWSProfile, region model.Region, endpoint *model.Endpoint) *session.Session {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable, // Ref. ~/.aws/config
		Profile:           profile.String(),
	}))

	session.Config.Region = aws.String(region.String())
	if endpoint != nil {
		// If you want to debug, uncomment the following lines.
		// session.Config.WithLogLevel(aws.LogDebugWithHTTPBody)
		session.Config.S3ForcePathStyle = aws.Bool(true)
		session.Config.Endpoint = aws.String(endpoint.String())
		session.Config.DisableSSL = aws.Bool(true)
	}
	return session
}

// detectContentType detects the content type of the file.
func detectContentType(reader io.Reader, filename string) (string, error) {
	// We determine CSS and JavaScript based on file extension
	// because we cannot determine them by mimetype.DetectReader.
	var extensionToContentType = map[string]string{
		".css": "text/css",
		".js":  "application/javascript",
	}
	contentType, found := extensionToContentType[filepath.Ext(filename)]
	if found {
		return contentType, nil
	}

	mtype, err := mimetype.DetectReader(reader)
	if err != nil {
		return "", errfmt.Wrap(service.ErrNotDetectContentType, err.Error())
	}
	return mtype.String(), nil
}

// duplicateReader duplicates the io.Reader.
func duplicateReader(r io.Reader) (io.Reader, io.Reader, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, err
	}
	return bytes.NewReader(data), bytes.NewReader(data), nil
}
