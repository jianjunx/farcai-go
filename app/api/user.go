package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"

	"github.com/gogf/gf/net/ghttp"
)

var User = userApi{}

type userApi struct{}

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

func (*userApi) Login(r *ghttp.Request) {
	var p *model.LoginReq
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
	}
	resp, err := service.User.Login(p)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{Code: 200, Data: resp})
}
