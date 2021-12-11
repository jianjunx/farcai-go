package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"

	"github.com/gogf/gf/net/ghttp"
)

var Assets = &assetsApi{}

type assetsApi struct{}

// 获取腾讯云临时密钥
func (*assetsApi) COSCredentials(r *ghttp.Request) {
	cred, err := service.Assets.COSCredentials()
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(&model.Response{
		Code: 200,
		Data: cred,
	})
}
