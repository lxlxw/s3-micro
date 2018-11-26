package server

import (
	"golang.org/x/net/context"

	"wps_store/service"

	pb "wps_store/rpc"
)

func (h apiService) PutObject(ctx context.Context, r *pb.PutObjectRequest) (*pb.PutObjectResponse, error) {
	res := service.PutObject(r)
	return &res, nil
}

func (h apiService) GetObject(ctx context.Context, r *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	res := service.GetObject(r)
	return &res, nil
}

func (h apiService) GetObjectPresignedUrl(ctx context.Context, r *pb.GetObjectPresignedUrlRequest) (*pb.GetObjectPresignedUrlResponse, error) {
	res := service.GetObjectPresignedUrl(r)
	return &res, nil
}

func (h apiService) PutObjectPresignedUrl(ctx context.Context, r *pb.PutObjectPresignedUrlRequest) (*pb.PutObjectPresignedUrlResponse, error) {
	res := service.PutObjectPresignedUrl(r)
	return &res, nil
}
