package s3

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
	"github.com/ks3sdklib/aws-sdk-go/service/s3/s3iface"
)

// Object has object body and metadata
type Object struct {
	Body        []byte
	ContentType string
}

// GetObject download a file on S3
func GetObject(svc s3iface.S3API, path Path) Object {
	input := &s3.GetObjectInput{
		Bucket: aws.String(path.Bucket),
		Key:    aws.String(path.Key),
	}
	res, err := svc.GetObject(input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(res.Body); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return Object{
		Body:        buf.Bytes(),
		ContentType: *res.ContentType,
	}
}

// PutObject upload a file to S3
func (svc *S3) PutObject() error {

	input := &s3.PutObjectInput{
		Bucket:             aws.String("wpspdftools"), // bucket名称
		Key:                aws.String("/hahatest/xxxxxxxxxxhhhh.txt"),  // object key
		ACL:                aws.String("public-read"),//权限，支持private(私有)，public-read(公开读)
		Body:               bytes.NewReader([]byte("PAYLOAD")),//要上传的内容
		ContentType:        aws.String("application/ocet-stream"),//设置content-type
	}
	if _, err := svc.S3.PutObject(input); err != nil {
		return err
	}
	return nil
}