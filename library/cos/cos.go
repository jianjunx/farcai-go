package cos

import (
	"os"
	"time"

	"github.com/gogf/gf/frame/g"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

var (
	Option    *sts.CredentialOptions
	bucket    string
	appid     string
	region    string
	path      string
	secretID  string
	secretKey string
)

func init() {
	secretID = os.Getenv("COS_SECRETID")
	secretKey = os.Getenv("COS_SECRETKEY")
	region = g.Cfg().GetString("cos.Region")
	bucket = g.Cfg().GetString("cos.Bucket")
	appid = g.Cfg().GetString("cos.Appid")
	path = g.Cfg().GetString("cos.Path")

	Option = getOpt()
}

func getOpt() *sts.CredentialOptions {
	return &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					// 密钥的权限列表。简单上传和分片需要以下的权限，其他权限列表请看 https://cloud.tencent.com/document/product/436/31923
					Action: []string{
						// 简单上传
						"name/cos:PostObject",
						"name/cos:PutObject",
					},
					Effect: "allow",
					Resource: []string{
						//这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + path + "/*",
					},
				},
			},
		},
	}
}

func GetClient() *sts.Client {
	client := sts.NewClient(
		secretID,
		secretKey,
		nil,
	)

	return client
}
