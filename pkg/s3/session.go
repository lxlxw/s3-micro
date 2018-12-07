package s3

import (
	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/aws/credentials"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"

	"github.com/lxlxw/s3-micro/pkg/util"
)

// S3 is a client for Amazon S3.
type S3 struct {
	*s3.S3
}

// S3 config info
type S3Conf struct {
	// S3 access key
	AccessKey string

	// S3 secret key
	Secretkey string

	// S3 region
	Region string

	// S3 endpoint
	Endpoint string
}

// New returns a new S3 client.
func New() (*S3, error) {
	conf := GetS3Conf()

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

// Get S3 config info.
func GetS3Conf() S3Conf {
	util.SetConfig(S3_CONF_FILENNAME)
	s3 := util.GetConfig().S3

	s3Conf := S3Conf{
		AccessKey: s3.AccessKey,
		Secretkey: s3.Secretkey,
		Region:    s3.Region,
		Endpoint:  s3.Endpoint,
	}
	return s3Conf
}
