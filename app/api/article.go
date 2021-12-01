package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"

	"github.com/gogf/gf/net/ghttp"
)

var Article = &articleApi{}

type articleApi struct{}

func (*articleApi) AddArticle(r *ghttp.Request) {
	var p *model.ArticleItem
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
		return
	}
	p.UserName = r.GetCtxVar("UserName").String()
	err := service.Article.AddArticle(p)
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
		Data: p,
	})
}

func (*articleApi) UpdateArticle(r *ghttp.Request) {
	var p *model.ArticleItem
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
		return
	}
	err := service.Article.UpdateArticle(p)
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
	})
}

func (*articleApi) DeleteArticle(r *ghttp.Request) {
	err := service.Article.DeleteArticle(r.GetInt64("id"))
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
	})
}
