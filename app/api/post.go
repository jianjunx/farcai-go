package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"

	"github.com/gogf/gf/net/ghttp"
)

var Post = &postApi{}

type postApi struct{}

// 添加文章
func (*postApi) AddPost(r *ghttp.Request) {
	var p *model.PostReq
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
	}
	p.UserId = r.GetCtxVar("uid").Int()
	res, err := service.Post.AddPost(p)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	pid, err := res.LastInsertId()
	if err != nil {
		service.ErrorHandler(r, err)
	}
	p.Pid = int(pid)
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
		Data: p,
	})
}

// 更新文章
func (*postApi) UpdatePost(r *ghttp.Request) {
	var p *model.PostReq
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
	}
	err := service.Post.UpdatePost(p)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
	})
}

// 获取文章详情
func (*postApi) GetPost(r *ghttp.Request) {
	pid := r.GetInt("id")
	post, err := service.Post.GetPostItem(&pid)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
		Data: post,
	})
}

// 删除文章
func (*postApi) DeletePost(r *ghttp.Request) {
	pid := r.GetInt("id")
	err := service.Post.DeletePost(&pid)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
	})
}

// 搜索文章
func (*postApi) SearchPost(r *ghttp.Request) {
	search := r.GetQueryString("val")
	posts, err := service.Post.SearchPost(&search)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
		Data: posts,
	})
}
