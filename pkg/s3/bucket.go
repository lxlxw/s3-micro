package s3

import (
	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
)

// Creates a new bucket.
func (svc *S3) CreateBucket(bucketName, acl string) error {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
		ACL:    aws.String(acl),
	}
	if _, err := svc.S3.CreateBucket(input); err != nil {
		return err
	}
	return nil
}

func (svc *S3) DeleteBucket(bucketName string) error {
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(bucketName),
	}
	if _, err := svc.S3.DeleteBucket(input); err != nil {
		return err
	}
	return nil
}

func (svc *S3) ListBucket() (*s3.ListBucketsOutput, error) {
	input := &s3.ListBucketsInput{}
	res, err := svc.S3.ListBuckets(input)
	if err != nil {
		return res, err
	}
	return res, nil
}
