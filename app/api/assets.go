package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"
	"os"

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

// 请求触发备份
func (*assetsApi) BackupDB(r *ghttp.Request) {
	// 验证下secret
	if r.GetString("secret") != os.Getenv("JWT_SECRET") {
		r.Response.WriteJsonExit(&model.Response{Code: 403, Data: "验证权限失败"})
		return
	}
	go service.Assets.BackupCategory()
	go service.Assets.BackupPost()
	go service.Assets.BackupUser()

	r.Response.WriteJsonExit(&model.Response{Code: 200, Data: "执行成功"})
}
