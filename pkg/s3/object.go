package s3

import (
	"bufio"
	"net/url"
	"strings"
	"time"

	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
)

// Retrieves objects from Amazon S3.
func (svc *S3) GetObject(bucketName, key, contentType string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket:              aws.String(bucketName),
		Key:                 aws.String(key),
		ResponseContentType: aws.String(contentType),
	}
	res, err := svc.S3.GetObject(input)
	if err != nil {
		return "", err
	}
	br := bufio.NewReader(res.Body)
	resBody, _ := br.ReadString('\n')
	return resBody, nil
}

// The HEAD operation without returning the object itself.
func (svc *S3) HeadObject(bucketName, key string) (bool, error) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	if _, err := svc.S3.HeadObject(input); err != nil {
		return false, err
	}
	return true, nil
}

// The HEAD operation returning the object presigned url.
func (svc *S3) HeadObjectPresignedUrl(bucketName, key string, expireTime int64) (*url.URL, bool) {
	input := &s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	res, err := svc.S3.HeadObjectPresignedUrl(input, time.Duration(expireTime*1000*1000*1000))
	if err != nil {
		return res, false
	}
	return res, true
}

// Puts object
func (svc *S3) PutObject(bucketName, key, fileContent, contentType, publicRead string, size, expireTime int64) error {
	input := &s3.PutObjectInput{
		Bucket:           aws.String(bucketName),         // bucket name
		Key:              aws.String(key),                // object key
		ACL:              aws.String(publicRead),         // acl, private, public-read
		Body:             strings.NewReader(fileContent), // bytes.NewReader([]byte(fileContent)),
		ContentType:      aws.String(contentType),        // set content-type
		ContentMaxLength: aws.Long(size),
		Expires:          aws.Time(time.Now().Add(time.Second * time.Duration(expireTime))),
	}
	if _, err := svc.S3.PutObject(input); err != nil {
		return err
	}
	return nil
}

// Gets object returning the object presigned url.
func (svc *S3) GetObjectPresignedUrl(bucketName, key string, expireTime int64) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}
	resp, err := svc.S3.GetObjectPresignedUrl(input, time.Duration(expireTime*1000*1000*1000))
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

// Puts object returning the object presigned url.
func (svc *S3) PutObjectPresignedUrl(bucketName, key, contentType, publicRead string, size, expireTime int64) (string, error) {
	input := &s3.PutObjectInput{
		Bucket:           aws.String(bucketName),
		Key:              aws.String(key),
		ACL:              aws.String(publicRead),
		ContentType:      aws.String(contentType),
		ContentMaxLength: aws.Long(size),
	}
	resp, err := svc.S3.PutObjectPresignedUrl(input, time.Duration(expireTime*1000*1000*1000))
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}
