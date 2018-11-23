package service

import (
	"log"
	//"bytes"

	//"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/aws/awserr"
	//"github.com/ks3sdklib/aws-sdk-go/service/s3"

	ws3 "wps_store/pkg/s3"

	pb "wps_store/proto"
)

/*
* 上传文件
*
 */
func PutObject(Name string) pb.PutObjectResponse {

	var result pb.PutObjectResponse

	//判断是ks3还是s3，去配置文件找出相应配置
	//New封装
	//params传入
	//如果带id，则先查找下该文件是否还存在，如果存在，直接返回外链

	client, err := ws3.New()
	if err != nil {
	}

	err1 := client.PutObject()
	if reqErr, ok := err1.(awserr.RequestFailure); ok {
		return pb.PutObjectResponse{Msg: err.(awserr.Error).Code(), Code: int32(reqErr.StatusCode())}
	}

	log.Println("sss")
	
	result = pb.PutObjectResponse{
		Msg:  "success",
		Code: 200,
		Data: Name,
	}
	return result
}
