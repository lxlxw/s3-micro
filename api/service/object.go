package service

import (
	"github.com/lxlxw/s3-micro/pkg/s3"

	pb "github.com/lxlxw/s3-micro/proto"
)

// Puts s3 object
func PutObject(r *pb.PutObjectRequest) pb.PutObjectResponse {
	client, _ := s3.New()

	objectExist, exErr := client.HeadObject(r.Bucketname, r.Key)
	if exErr != nil {
		return pb.PutObjectResponse{Msg: exErr.Error(), Code: 403}
	}
	if objectExist == true {
		url, getErr := client.GetObjectPresignedUrl(r.Bucketname, r.Key, r.Expiretime)
		if getErr != nil {
			return pb.PutObjectResponse{Msg: getErr.Error(), Code: 403}
		}
		return pb.PutObjectResponse{Msg: "success", Code: 200, Data: url}
	}
	putErr := client.PutObject(r.Bucketname, r.Key, r.Filecontent, r.Contenttype, r.Publicread, r.Contentmaxlength, r.Expiretime)
	if putErr != nil {
		return pb.PutObjectResponse{Msg: putErr.Error(), Code: 403}
	}

	url, getpErr := client.GetObjectPresignedUrl(r.Bucketname, r.Key, r.Expiretime)
	if getpErr != nil {
		return pb.PutObjectResponse{Msg: getpErr.Error(), Code: 403}
	}
	return pb.PutObjectResponse{Msg: "success", Code: 200, Data: url}
}

// Downloads s3 object
func GetObject(r *pb.GetObjectRequest) pb.GetObjectResponse {
	client, _ := s3.New()

	resBody, err := client.GetObject(r.Bucketname, r.Key, r.Contenttype)
	if err != nil {
		return pb.GetObjectResponse{Msg: err.Error(), Code: 403}
	}
	return pb.GetObjectResponse{Msg: "success", Code: 200, Data: resBody}
}

// Gets s3 object presigned url
func GetObjectPresignedUrl(r *pb.GetObjectPresignedUrlRequest) pb.GetObjectPresignedUrlResponse {
	client, _ := s3.New()

	url, err := client.GetObjectPresignedUrl(r.Bucketname, r.Key, r.Expiretime)
	if err != nil {
		return pb.GetObjectPresignedUrlResponse{Msg: err.Error(), Code: 403}
	}
	return pb.GetObjectPresignedUrlResponse{Msg: "success", Code: 200, Data: url}
}

// Puts s3 object presigned url
func PutObjectPresignedUrl(r *pb.PutObjectPresignedUrlRequest) pb.PutObjectPresignedUrlResponse {
	client, _ := s3.New()

	presignedUrl, err := client.PutObjectPresignedUrl(r.Bucketname, r.Key, r.Contenttype, r.Publicread, r.Contentmaxlength, r.Expiretime)
	if err != nil {
		return pb.PutObjectPresignedUrlResponse{Msg: err.Error(), Code: 403}
	}
	return pb.PutObjectPresignedUrlResponse{Msg: "success", Code: 200, Data: presignedUrl}
}
