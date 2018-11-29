package controller

import (
	"golang.org/x/net/context"

	"wps_store/api/service"

	pb "wps_store/rpc"
)

func (h ApiService) CreateBucket(ctx context.Context, r *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	res := service.CreateBucket(r)
	return &res, nil
}
