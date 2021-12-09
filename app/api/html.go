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

// 首页
func (*htmlApi) Home(r *ghttp.Request) {
	page := r.GetQueryInt("page", 1)
	ctgId := r.GetQueryInt("cid")
	posts, categorys, total, err := service.Html.Home(&ctgId, &page)
	if err != nil {
		service.ErrorHandler(r, err)
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

// 文章详情页
func (*htmlApi) Detail(r *ghttp.Request) {
	ids := strings.Split(r.GetString("id"), ".")
	if len(ids) == 0 {
		service.ErrorHandler(r, errors.New("ID错误"))
	}
	id := gconv.Int(ids[0])
	article, err := service.Html.Detail(&id)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteTpl("layout.html", g.Map{
		"main":    "detail.html",
		"title":   article.Title,
		"article": article,
	})
	// 累计查看次数
	go service.Html.AddViewCount(&article.Pid)
}

// 编辑器 写作页
func (*htmlApi) Writing(r *ghttp.Request) {
	categorys, err := service.Category.GetCategorys()
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteTpl("writing.html", g.Map{
		"categorys": categorys,
		"cosBucket": g.Cfg().GetString("cos.Bucket"),
		"cosRegion": g.Cfg().GetString("cos.Region"),
		"cosPath":   g.Cfg().GetString("cos.Path"),
	})
}

// 登录页
func (*htmlApi) Login(r *ghttp.Request) {
	r.Response.WriteTpl("layout.html", g.Map{
		"main":  "login.html",
		"title": "登录",
	})
}

// 归档页面
func (*htmlApi) Pigeonhole(r *ghttp.Request) {
	lines, categorys, err := service.Html.Pigeonhole()
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteTpl("layout.html", g.Map{
		"main":      "pigeonhole.html",
		"title":     "归档",
		"categorys": categorys,
		"lines":     lines,
	})
}

// 分类页
func (*htmlApi) Category(r *ghttp.Request) {
	page := r.GetQueryInt("page", 1)
	ctgId := r.GetInt("cid", 0)
	if ctgId == 0 {
		r.Response.RedirectTo("/")
	}
	posts, categorys, total, err := service.Html.Home(&ctgId, &page)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	categoryName := ""
	if len(*posts) > 0 {
		// 取当前分类名
		_p := *posts
		categoryName = _p[0].CategoryName
	}
	pages := service.Html.GetPages(float64(total))
	r.Response.WriteTpl("layout.html", g.Map{
		"main":         "category.html",
		"title":        categoryName, // 副标题
		"posts":        posts,
		"categorys":    categorys,
		"categoryName": categoryName,
		"total":        total,
		"page":         page,
		"pageEnd":      int(page) != len(pages),
		"pages":        pages,
	})
}

// 自定义页面
func (*htmlApi) CustomPage(r *ghttp.Request) {
	custom := r.GetString("custom")
	empt, post, err := service.Html.Custom(&custom)
	if empt {
		r.Response.RedirectTo("/")
	}
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteTpl("layout.html", g.Map{
		"main":    "custom.html",
		"title":   post.Title, // 副标题
		"article": post,
	})
	// 累计查看次数
	go service.Html.AddViewCount(&post.Pid)
}
