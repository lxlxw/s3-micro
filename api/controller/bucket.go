package controller

import (
	"golang.org/x/net/context"

	"github.com/lxlxw/s3-micro/api/service"

	pb "github.com/lxlxw/s3-micro/proto"
)

// Creates a new S3 bucket.
func (h ApiService) CreateBucket(ctx context.Context, r *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	res := service.CreateBucket(r)
	return &res, nil
}

func (h ApiService) DeleteBucket(ctx context.Context, r *pb.DeleteBucketRequest) (*pb.DeleteBucketResponse, error) {
	res := service.DeleteBucket(r)
	return &res, nil
}

func (h ApiService) ListBucket(ctx context.Context, r *pb.ListBucketRequest) (*pb.ListBucketResponse, error) {
	res := service.ListBucket(r)
	return &res, nil
}
