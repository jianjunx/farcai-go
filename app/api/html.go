package api

import (
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
