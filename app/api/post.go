package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"

	"github.com/gogf/gf/net/ghttp"
)

var Post = &postApi{}

type postApi struct{}

func (*postApi) AddPost(r *ghttp.Request) {
	var p *model.PostReq
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
		return
	}
	p.UserId = r.GetCtxVar("uid").Int()
	res, err := service.Post.AddPost(p)
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	pid, err := res.LastInsertId()
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	p.Pid = int(pid)
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
		Data: p,
	})
}

func (*postApi) UpdatePost(r *ghttp.Request) {
	var p *model.PostReq
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
		return
	}
	err := service.Post.UpdatePost(p)
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
	})
}

func (*postApi) GetPost(r *ghttp.Request) {
	pid := r.GetInt("id")
	post, err := service.Post.GetPostItem(&pid)
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
		Data: post,
	})
}

func (*postApi) DeletePost(r *ghttp.Request) {
	pid := r.GetInt("id")
	err := service.Post.DeletePost(&pid)
	if err != nil {
		service.ErrorHandler(r, err)
		return
	}
	r.Response.WriteJsonExit(model.Response{
		Code: 200,
	})
}
