package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "wps_store/rpc"
)

var (
	ServerPort string
	EndPoint   string
)

type apiService struct {
	ssss string
}

func NewStoreApiService() *apiService {
	return &apiService{}
}

func ErrorResponse() (*pb.PutObjectResponse, error) {
	return &pb.PutObjectResponse{Msg: "操作失败!", Code: -1, Data: ""}, nil
}

func Run() (err error) {
	EndPoint = ":" + ServerPort

	lis, err := net.Listen("tcp", EndPoint)
	log.Println("Listen success:", EndPoint)
	if err != nil {
		log.Fatalf("Grpc服务启动失败: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStoreApiServiceServer(s, NewStoreApiService())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Grpc服务启动失败: %v", err)
	}
	return err
}
