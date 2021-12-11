package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var User = userApi{}

type userApi struct{}

// 注册页面
func (*userApi) Register(r *ghttp.Request) {
	var p *model.RegisterReq
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
	}
	err := service.User.Register(p)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{Code: 200})
}

// 用户登录
func (*userApi) Login(r *ghttp.Request) {
	var p *model.LoginReq
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
	}
	ip := r.GetClientIp()
	resp, err := service.User.Login(p, &ip)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{Code: 200, Data: resp})
}

// 获取用户信息 暂时没用
func (*userApi) GetUserInfo(r *ghttp.Request) {
	uid := gconv.Int(r.GetCtxVar("uid"))
	user, err := service.User.GetUserInfo(&uid)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
		Data: user,
	})
}
