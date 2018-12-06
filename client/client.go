package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/lxlxw/s3-micro/proto"
)

func main() {

	// conn
	conn, err := grpc.Dial(":50052", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	// new client
	c := pb.NewStoreApiServiceClient(conn)

	// call method
	req := &pb.GetConfigStoreInfoRequest{Store: "ks3"}
	res, err := c.GetConfigStoreInfo(context.Background(), req)

	if err != nil {
		log.Println(err)
	}

	log.Println(res.Data)

}
