package s3

import (
	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/aws/credentials"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"

	"wps_store/pkg/util"
)

type S3 struct {
	*s3.S3
}

type S3Conf struct {
	AccessKey string
	Secretkey string
	Region    string
	Endpoint  string
}

func New(store string) (*S3, error) {
	conf := GetS3Conf(store)

	credentials := credentials.NewStaticCredentials(conf.AccessKey, conf.Secretkey, "")
	client := s3.New(&aws.Config{
		Region:           conf.Region,
		Credentials:      credentials,
		Endpoint:         conf.Endpoint,
		DisableSSL:       true,  //是否禁用https
		LogLevel:         1,     //是否开启日志,0为关闭日志，1为开启日志
		S3ForcePathStyle: false, //是否强制使用path style方式访问
		LogHTTPBody:      true,  //是否把HTTP请求body打入日志
	})
	return &S3{S3: client}, nil
}

func GetS3Conf(store string) S3Conf {
	util.SetConfig()
	s3Conf := S3Conf{}
	ks3 := util.GetConfig().Ks3
	as3 := util.GetConfig().As3

	switch store {
	case "ks3":
		s3Conf = S3Conf{AccessKey: ks3.AccessKey, Secretkey: ks3.Secretkey, Region: ks3.Region, Endpoint: ks3.Endpoint}
	case "as3":
		s3Conf = S3Conf{AccessKey: as3.AccessKey, Secretkey: as3.Secretkey, Region: as3.Region, Endpoint: as3.Endpoint}
	default:
		s3Conf = S3Conf{AccessKey: ks3.AccessKey, Secretkey: ks3.Secretkey, Region: ks3.Region, Endpoint: ks3.Endpoint}
	}
	return s3Conf
}
