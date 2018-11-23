package s3

import (

	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/aws/credentials"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
)

// S3 embeds *s3.S3 to be used to call New
type S3 struct {
	*s3.S3
}

func New() (*S3, error) {
	credentials := credentials.NewStaticCredentials("AKLTRPycUPrRSDOP492EPQO6Bw","OCq39hfC3rUXyUsmapV7X740zoFHs++ZQGP97MZJZgjqQ1EnEA6lDQAy3XyWkd8a1Q==","")
	client := s3.New(&aws.Config{
		Region: "BEIJING",
		Credentials: credentials,
		Endpoint:"ks3-cn-beijing.ksyun.com",//s3地址
		DisableSSL:true,//是否禁用https
		LogLevel:1,//是否开启日志,0为关闭日志，1为开启日志
		S3ForcePathStyle:false,//是否强制使用path style方式访问
		LogHTTPBody:true,//是否把HTTP请求body打入日志
		})
	return &S3{S3: client}, nil
}