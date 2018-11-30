package rpcserver

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"wps_store/api/controller"
	"wps_store/api/middleware"
	pb "wps_store/rpc"
)

var (
	ServerPort string
	EndPoint   string
)

type apiService struct {
	controller.ApiService
}

func NewStoreApiService() *apiService {
	return &apiService{}
}

func ErrorResponse() (*pb.PutObjectResponse, error) {
	return &pb.PutObjectResponse{Msg: "操作失败!", Code: -1, Data: ""}, nil
}

func RunServer() (err error) {
	EndPoint = ":" + ServerPort

	lis, err := net.Listen("tcp", EndPoint)
	log.Println("Listen success:", EndPoint)
	if err != nil {
		log.Fatalf("Grpc服务启动失败: %v", err)
	}
	// 注册interceptor
	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptorChain(middleware.Recovery, middleware.Logging)))
	pb.RegisterStoreApiServiceServer(s, NewStoreApiService())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Grpc服务启动失败: %v", err)
	}
	return err
}

func UnaryInterceptorChain(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		chain := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			chain = build(interceptors[i], chain, info)
		}
		return chain(ctx, req)
	}
}

func build(c grpc.UnaryServerInterceptor, n grpc.UnaryHandler, info *grpc.UnaryServerInfo) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return c(ctx, req, info, n)
	}
}
