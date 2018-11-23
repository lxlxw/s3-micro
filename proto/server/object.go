package server

import (
	"log"
    "golang.org/x/net/context"
	
	"wps_store/service"

	pb "wps_store/proto"
)

func (h apiService) PutObject(ctx context.Context, r *pb.PutObjectRequest) (*pb.PutObjectResponse, error) {

	if r.Bucket != "" {
		log.Println("获取数据:", r.Bucket)
		res := service.PutObject(r.Bucket)
		return &res, nil
	}
	return ErrorResponse()
}

func (h apiService) GetObject(ctx context.Context, r *pb.GetObjectRequest) (*pb.GetObjectResponse, error) {
	return &pb.GetObjectResponse{Res: "操作失败!"}, nil
}