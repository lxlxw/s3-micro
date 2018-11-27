package server

import (
	"golang.org/x/net/context"

	"wps_store/service"

	pb "wps_store/rpc"
)

func (h apiService) CreateBucket(ctx context.Context, r *pb.CreateBucketRequest) (*pb.CreateBucketResponse, error) {
	res := service.CreateBucket(r)
	return &res, nil
}
