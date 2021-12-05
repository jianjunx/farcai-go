package api

import (
	"errors"
	"farcai-go/app/service"
	"strings"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

var Html = &htmlApi{}

type htmlApi struct{}

func (*htmlApi) Home(r *ghttp.Request) {
	page := r.GetQueryInt("page", 1)
	ctgId := r.GetQueryInt("cid")
	posts, categorys, total, err := service.Html.Home(&ctgId, &page)
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	pages := service.Html.GetPages(float64(total))
	r.Response.WriteTpl("layout.html", g.Map{
		"main":      "home.html",
		"posts":     posts,
		"categorys": categorys,
		"total":     total,
		"page":      page,
		"pageEnd":   int(page) != len(pages),
		"pages":     pages,
	})
}

func (*htmlApi) Detail(r *ghttp.Request) {
	ids := strings.Split(r.GetString("id"), ".")
	if len(ids) == 0 {
		service.ErrorHandler(r, errors.New("ID错误"))
		return
	}
	id := gconv.Int(ids[0])
	article, err := service.Html.Detail(&id)
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}

	r.Response.WriteTpl("layout.html", g.Map{
		"main":    "detail.html",
		"title":   article.Title,
		"article": article,
	})
	// 累计查看次数
	go service.Html.AddViewCount(&article.Pid)
}

func (*htmlApi) Writing(r *ghttp.Request) {
	categorys, err := service.Category.GetCategorys()
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteTpl("writing.html", g.Map{
		"id":        1,
		"categorys": categorys,
		"cosBucket": g.Cfg().GetString("cos.Bucket"),
		"cosRegion": g.Cfg().GetString("cos.Region"),
		"cosPath":   g.Cfg().GetString("cos.Path"),
	})
}

func (*htmlApi) Login(r *ghttp.Request) {
	r.Response.WriteTpl("layout.html", g.Map{
		"main":  "login.html",
		"title": "登录",
	})
}
