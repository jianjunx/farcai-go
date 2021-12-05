package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"

	"github.com/gogf/gf/net/ghttp"
)

var Assets = &assetsApi{}

type assetsApi struct{}

func (*assetsApi) COSCredentials(r *ghttp.Request) {
	cred, err := service.Assets.COSCredentials()
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteJsonExit(&model.Response{
		Code: 200,
		Data: cred,
	})
}
