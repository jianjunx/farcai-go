package cos

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gogf/gf/frame/g"
	sdk "github.com/tencentyun/cos-go-sdk-v5"
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

// 配置项
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

func getSdkClient() *sdk.Client {
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region))
	b := &sdk.BaseURL{BucketURL: u}
	return sdk.NewClient(b, &http.Client{
		Transport: &sdk.AuthorizationTransport{
			SecretID:  secretID,  // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: secretKey, // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
}

// 通过字符串上传对象到COS
func StringUpload(name, text *string) (*sdk.Response, error) {
	f := strings.NewReader(*text)
	return getSdkClient().Object.Put(context.Background(), *name, f, nil)
}
