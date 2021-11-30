package api

import (
	"farcai-go/app/service"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Html = &htmlApi{}

type htmlApi struct{}

func (*htmlApi) Home(r *ghttp.Request) {
	r.Response.WriteTpl("layout.html", g.Map{
		"main": "home.html",
		"id":   112,
		"name": "JJX",
	})
}

func (*htmlApi) Writing(r *ghttp.Request) {
	categorys, err := service.Category.GetCategorys()
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteTpl("writing.html", g.Map{"id": 1, "categorys": categorys})
}

func (*htmlApi) Login(r *ghttp.Request) {
	r.Response.WriteTpl("layout.html", g.Map{
		"main": "login.html",
		"title": "登录",
	})
}