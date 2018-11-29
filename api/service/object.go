package service

import (
	ws3 "wps_store/pkg/s3"

	pb "wps_store/rpc"
)

/*
* 上传文件
*
 */
func PutObject(r *pb.PutObjectRequest) pb.PutObjectResponse {
	client, _ := ws3.New(r.Store)

	//判断文件是否存在
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

/*
*  下载文件
*
 */
func GetObject(r *pb.GetObjectRequest) pb.GetObjectResponse {
	client, _ := ws3.New(r.Store)

	resBody, err := client.GetObject(r.Bucketname, r.Key, r.Contenttype)
	if err != nil {
		return pb.GetObjectResponse{Msg: err.Error(), Code: 403}
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
	if err != nil {
		return pb.GetObjectPresignedUrlResponse{Msg: err.Error(), Code: 403}
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
	if err != nil {
		return pb.PutObjectPresignedUrlResponse{Msg: err.Error(), Code: 403}
	}
	return pb.PutObjectPresignedUrlResponse{Msg: "success", Code: 200, Data: presignedUrl}
}
