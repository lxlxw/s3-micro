package s3

import (
	"io"
	"reflect"
	"testing"

	"github.com/ks3sdklib/aws-sdk-go/service/s3"
	"github.com/ks3sdklib/aws-sdk-go/service/s3/s3iface"
)

type mockedGetObject struct {
	s3iface.S3API
	Resp Object
}

func (m *mockedGetObject) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	pr, pw := io.Pipe()

	go func() {
		pw.Write(m.Resp.Body)
		pw.Close()
	}()

	return &s3.GetObjectOutput{
		Body:        pr,
		ContentType: &m.Resp.ContentType,
	}, nil
}

func TestGetObjectSuccess(t *testing.T) {
	want := Object{
		Body:        []byte("body"),
		ContentType: "text/plain",
	}

	svc := &mockedGetObject{Resp: want}
	path := Path{Bucket: "bucket", Key: "key"}

	got := GetObject(svc, path)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Want %q, got %q", want, got)
	}
}