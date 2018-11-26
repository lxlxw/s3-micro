package service

import (
	"github.com/ks3sdklib/aws-sdk-go/aws/awserr"

	ws3 "wps_store/pkg/s3"

	pb "wps_store/rpc"
)

/*
* 上传文件
*
 */
func PutObject(r *pb.PutObjectRequest) pb.PutObjectResponse {
	client, _ := ws3.New(r.Store)

	//TODO: 要处理下key， 拆分. 然后将keyname做成文件目录

	//判断文件是否存在
	objectExist := client.HeadObject(r.Bucketname, r.Key)
	if objectExist == true {
		url, err := client.GetObjectPresignedUrl(r.Bucketname, r.Key, r.Expiretime)
		if err == nil {
			return pb.PutObjectResponse{Msg: "success", Code: 200, Data: url}
		}
	}

	err := client.PutObject(r.Bucketname, r.Key, r.Filecontent, r.Contenttype, r.Publicread, r.Contentmaxlength, r.Expiretime)
	if reqErr, ok := err.(awserr.RequestFailure); ok {
		return pb.PutObjectResponse{Msg: err.(awserr.Error).Code(), Code: int32(reqErr.StatusCode())}
	}
	url, err := client.GetObjectPresignedUrl(r.Bucketname, r.Key, r.Expiretime)
	if err != nil {
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			return pb.PutObjectResponse{Msg: err.(awserr.Error).Code(), Code: int32(reqErr.StatusCode())}
		}
	}
	return pb.PutObjectResponse{Msg: "success", Code: 200, Data: url}
}

/*
*  下载文件
*
 */
func GetObject(r *pb.GetObjectRequest) pb.GetObjectResponse {
	client, _ := ws3.New(r.Store)

	resBody, err := client.GetObject(r.Bucketname, r.Key, r.Contenttype)
	if reqErr, ok := err.(awserr.RequestFailure); ok {
		return pb.GetObjectResponse{Msg: err.(awserr.Error).Code(), Code: int32(reqErr.StatusCode())}
	}
	return pb.GetObjectResponse{Msg: "success", Code: 200, Data: resBody}
}

/*
*  获取文件下载外链
*
 */
func GetObjectPresignedUrl(r *pb.GetObjectPresignedUrlRequest) pb.GetObjectPresignedUrlResponse {
	client, _ := ws3.New(r.Store)

	url, err := client.GetObjectPresignedUrl(r.Bucketname, r.Key, r.Expiretime)
	if reqErr, ok := err.(awserr.RequestFailure); ok {
		return pb.GetObjectPresignedUrlResponse{Msg: err.(awserr.Error).Code(), Code: int32(reqErr.StatusCode())}
	}
	return pb.GetObjectPresignedUrlResponse{Msg: "success", Code: 200, Data: url}
}

/*
*  生成文件上传外链
*
 */
func PutObjectPresignedUrl(r *pb.PutObjectPresignedUrlRequest) pb.PutObjectPresignedUrlResponse {
	client, _ := ws3.New(r.Store)

	presignedUrl, err := client.PutObjectPresignedUrl(r.Bucketname, r.Key, r.Contenttype, r.Publicread, r.Contentmaxlength, r.Expiretime)
	if reqErr, ok := err.(awserr.RequestFailure); ok {
		return pb.PutObjectPresignedUrlResponse{Msg: err.(awserr.Error).Code(), Code: int32(reqErr.StatusCode())}
	}
	return pb.PutObjectPresignedUrlResponse{Msg: "success", Code: 200, Data: presignedUrl}
}
