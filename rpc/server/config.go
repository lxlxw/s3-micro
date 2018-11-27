package server

import (
	"golang.org/x/net/context"

	"wps_store/service"

	pb "wps_store/rpc"
)

func (h apiService) UpdateConfigStoreInfo(ctx context.Context, r *pb.UpdateConfigStoreInfoRequest) (*pb.UpdateConfigStoreInfoResponse, error) {
	res := service.UpdateConfigStoreInfo(r)
	return &res, nil
}

func (h apiService) GetConfigStoreInfo(ctx context.Context, r *pb.GetConfigStoreInfoRequest) (*pb.GetConfigStoreInfoResponse, error) {
	res := service.GetConfigStoreInfo(r)
	return &res, nil
}
