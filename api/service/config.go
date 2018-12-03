package service

import (
	"wps_store/pkg/s3"
	"wps_store/pkg/util"
	pb "wps_store/rpc"
)

/*
* 更新配置文件
*
 */
func UpdateConfigStoreInfo(r *pb.UpdateConfigStoreInfoRequest) pb.UpdateConfigStoreInfoResponse {
	conf := getS3Conf(r)
	fileString, err := util.EncodeConfig(&conf)
	if err != nil {
		return pb.UpdateConfigStoreInfoResponse{Msg: err.Error(), Code: 403}
	}
	fileErr := util.UpdateConfigFile(fileString, s3.S3_CONF_FILENNAME)
	if fileErr != nil {
		return pb.UpdateConfigStoreInfoResponse{Msg: err.Error(), Code: 403}
	}
	return pb.UpdateConfigStoreInfoResponse{Msg: "success", Code: 200, Data: ""}
}

/*
*  获取配置文件信息
*
 */
func GetConfigStoreInfo(r *pb.GetConfigStoreInfoRequest) pb.GetConfigStoreInfoResponse {
	conf := s3.GetS3Conf(r.Store)

	confInfo := &pb.ConfigInfo{
		Accesskey: conf.AccessKey,
		Secretkey: conf.Secretkey,
		Region:    conf.Region,
		Endpoint:  conf.Endpoint,
	}
	return pb.GetConfigStoreInfoResponse{Msg: "success", Code: 200, Data: confInfo}
}

/*
*  获取需要更新的配置信息
*
 */
func getS3Conf(r *pb.UpdateConfigStoreInfoRequest) util.Config {

	Ks3Conf := util.Ks3Config{}
	As3Conf := util.As3Config{}

	switch r.Store {
	case s3.StoreAs3:
		ks3 := s3.GetS3Conf(s3.StoreKs3)
		Ks3Conf = util.Ks3Config{AccessKey: ks3.AccessKey, Secretkey: ks3.Secretkey, Region: ks3.Region, Endpoint: ks3.Endpoint}
		As3Conf = util.As3Config{AccessKey: r.Accesskey, Secretkey: r.Secretkey, Region: r.Region, Endpoint: r.Domain}
	case s3.StoreKs3:
		as3 := s3.GetS3Conf(s3.StoreAs3)
		As3Conf = util.As3Config{AccessKey: as3.AccessKey, Secretkey: as3.Secretkey, Region: as3.Region, Endpoint: as3.Endpoint}
		Ks3Conf = util.Ks3Config{AccessKey: r.Accesskey, Secretkey: r.Secretkey, Region: r.Region, Endpoint: r.Domain}
	default:
		//default update ks3
		as3 := s3.GetS3Conf(s3.StoreAs3)
		As3Conf = util.As3Config{AccessKey: as3.AccessKey, Secretkey: as3.Secretkey, Region: as3.Region, Endpoint: as3.Endpoint}
		Ks3Conf = util.Ks3Config{AccessKey: r.Accesskey, Secretkey: r.Secretkey, Region: r.Region, Endpoint: r.Domain}
	}
	return util.Config{Ks3: Ks3Conf, As3: As3Conf}
}
