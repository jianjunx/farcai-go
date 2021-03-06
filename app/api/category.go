package api

import (
	"farcai-go/app/model"
	"farcai-go/app/service"

	"github.com/gogf/gf/net/ghttp"
)

var Category categoryApi

type categoryApi struct{}

// 添加分类
func (*categoryApi) AddCategory(r *ghttp.Request) {
	p := model.Category{}
	if err := r.Parse(&p); err != nil {
		service.ErrorHandler(r, err)
	}
	_, err := service.Category.AddCategory(&p.Name)
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{Code: 200})
}

func (*categoryApi) GetCategorys(r *ghttp.Request) {
	categorys, err := service.Category.GetCategorys()
	if err != nil {
		service.ErrorHandler(r, err)
	}
	r.Response.WriteJsonExit(model.Response{Code: 200, Data: categorys})
}
