package service

import (
	"github.com/lxlxw/s3-micro/pkg/s3"

	pb "github.com/lxlxw/s3-micro/proto"
)

// Creates s3 bucket
func CreateBucket(r *pb.CreateBucketRequest) pb.CreateBucketResponse {

	client, _ := s3.New()

	err := client.CreateBucket(r.Bucketname, r.Publicread)
	if err != nil {
		return pb.CreateBucketResponse{Msg: err.Error(), Code: 403}
	}
	return pb.CreateBucketResponse{Msg: "success", Code: 200, Data: ""}
}

// Delete s3 bucket
func DeleteBucket(r *pb.DeleteBucketRequest) pb.DeleteBucketResponse {

	client, _ := s3.New(r.Store)

	err := client.DeleteBucket(r.Bucketname)
	if err != nil {
		return pb.DeleteBucketResponse{Msg: err.Error(), Code: 403}
	}
	return pb.DeleteBucketResponse{Msg: "success", Code: 200, Data: ""}
}

// List s3 bucket
func ListBucket(r *pb.ListBucketRequest) pb.ListBucketResponse {

	client, _ := s3.New(r.Store)

	res, err := client.ListBucket()
	if err != nil {
		return pb.ListBucketResponse{Msg: err.Error(), Code: 403}
	}

	bucketList := res.Buckets
	ListBucket := &pb.ListBucket{}
	for i := 0; i < len(bucketList); i++ {
		Bucket := *bucketList[i].Name
		ListBucket.Bucket = append(ListBucket.Bucket, Bucket)
	}
	return pb.ListBucketResponse{Msg: "success", Code: 200, Data: ListBucket}
}
