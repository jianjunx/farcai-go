package service

import (
	"farcai-go/library/cos"

	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
)

var Assets = &assetsService{}

type assetsService struct{}

func (*assetsService) COSCredentials() (*sts.CredentialResult, error) {
	return cos.GetClient().GetCredential(cos.Option)
}
