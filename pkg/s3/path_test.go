package s3

import (
	"errors"
	"testing"
)

func TestParsePathSuccess(t *testing.T) {
	var validCases = []struct {
		in   string
		want Path
	}{
		{"s3://bucket/key", Path{Bucket: "bucket", Key: "/key"}},
		{"s3://bucket/key1/key2", Path{Bucket: "bucket", Key: "/key1/key2"}},
	}

	for _, v := range validCases {
		got, _ := ParsePath(v.in)
		if got != v.want {
			t.Errorf("Reverse(%q) == %q, want %q", v.in, got, v.want)
		}
	}
}

func TestParsePathFailure(t *testing.T) {
	var invalidCases = []struct {
		in   string
		want error
	}{
		{"bucket/key", errors.New("invalid S3 Path")},
		{"s3://", errors.New("invalid S3 Path")},
		{"s3://bucket", errors.New("invalid S3 Path")},
		{"s3://bucket/", errors.New("invalid S3 Path")},
		{"s3://bucket/key/", errors.New("invalid S3 Path")},
		{"http://bucket/key", errors.New("invalid S3 Path")},
	}

	for _, v := range invalidCases {
		_, err := ParsePath(v.in)
		if err == nil {
			t.Errorf("Reverse(%q) == %q, want %q", v.in, err, v.want)
		}
	}
}