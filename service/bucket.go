package service

import (
	ws3 "wps_store/pkg/s3"

	pb "wps_store/rpc"
)

func CreateBucket(r *pb.CreateBucketRequest) pb.CreateBucketResponse {

	client, _ := ws3.New(r.Store)

	err := client.CreateBucket(r.Bucketname, r.Publicread)
	if err != nil {
		return pb.CreateBucketResponse{Msg: err.Error(), Code: 403}
	}
	return pb.CreateBucketResponse{Msg: "success", Code: 200, Data: ""}
}
