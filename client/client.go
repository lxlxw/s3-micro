package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "wps_store/proto"
)

func main() {

	// 连接
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewStoreApiServiceClient(conn)

	// 调用方法
	req := &pb.StoreApiRequest{Referer: "gRPC"}
	res, err := c.ChainUpload(context.Background(), req)

	if err != nil {
		log.Println(err)
	}

	log.Println(res.Message)

}
