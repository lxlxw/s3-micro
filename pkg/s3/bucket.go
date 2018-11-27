package s3

import (
	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
)

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
