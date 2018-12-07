package rpcserver

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/lxlxw/s3-micro/api/controller"
	"github.com/lxlxw/s3-micro/api/middleware"
	pb "github.com/lxlxw/s3-micro/proto"
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

// Run server
func RunServer() (err error) {
	EndPoint = ":" + ServerPort

	lis, err := net.Listen("tcp", EndPoint)
	log.Println("Listen success:", EndPoint)
	if err != nil {
		log.Fatalf("Grpc服务启动失败: %v", err)
	}
	// new interceptor
	s := grpc.NewServer(grpc.UnaryInterceptor(UnaryInterceptorChain(middleware.Recovery, middleware.Logging)))
	pb.RegisterStoreApiServiceServer(s, NewStoreApiService())
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Grpc服务启动失败: %v", err)
	}
	return err
}

// Unary interceptor chain
func UnaryInterceptorChain(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		chain := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			chain = build(interceptors[i], chain, info)
		}
		return chain(ctx, req)
	}
}

// build interceptors
func build(c grpc.UnaryServerInterceptor, n grpc.UnaryHandler, info *grpc.UnaryServerInfo) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return c(ctx, req, info, n)
	}
}
