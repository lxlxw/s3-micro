package s3

import (
	"errors"
	"net/url"
	"strings"
)

// A Path have S3 bucket and key
type Path struct {
	Bucket string
	Key    string
}

// ParsePath parse a file path of S3, and return Path struct
func ParsePath(path string) (Path, error) {
	parsedURL, err := url.ParseRequestURI(path)
	if err != nil {
		return Path{}, errors.New("invalid S3 Path")
	}

	if isInvalidS3Path(*parsedURL) {
		return Path{}, errors.New("invalid S3 Path")
	}

	return Path{
		Bucket: parsedURL.Host,
		Key:    parsedURL.Path,
	}, nil
}

func isInvalidS3Path(url url.URL) bool {
	return url.Host == "" || url.Path == "" || url.Path == "/" || url.Scheme != "s3" || strings.HasSuffix(url.Path, "/")
}