package service

import (
	"github.com/lxlxw/s3-micro/pkg/s3"
	"github.com/lxlxw/s3-micro/pkg/util"
	pb "github.com/lxlxw/s3-micro/proto"
)

// Updates s3 config info
func UpdateConfigStoreInfo(r *pb.UpdateConfigStoreInfoRequest) pb.UpdateConfigStoreInfoResponse {

	s3Conf := util.Config{
		S3: util.S3Config{
			AccessKey: r.Accesskey,
			Secretkey: r.Secretkey,
			Region:    r.Region,
			Endpoint:  r.Domain,
		},
	}

	fileString, err := util.EncodeConfig(&s3Conf)
	if err != nil {
		return pb.UpdateConfigStoreInfoResponse{Msg: err.Error(), Code: 403}
	}
	fileErr := util.UpdateConfigFile(fileString, s3.S3_CONF_FILENNAME)
	if fileErr != nil {
		return pb.UpdateConfigStoreInfoResponse{Msg: err.Error(), Code: 403}
	}
	return pb.UpdateConfigStoreInfoResponse{Msg: "success", Code: 200, Data: ""}
}

// Gets s3 config info
func GetConfigStoreInfo(r *pb.GetConfigStoreInfoRequest) pb.GetConfigStoreInfoResponse {
	conf := s3.GetS3Conf()

	confInfo := &pb.ConfigInfo{
		Accesskey: conf.AccessKey,
		Secretkey: conf.Secretkey,
		Region:    conf.Region,
		Endpoint:  conf.Endpoint,
	}
	return pb.GetConfigStoreInfoResponse{Msg: "success", Code: 200, Data: confInfo}
}
