package service

import (
	"farcai-go/app/model"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
)



func ErrorHandler(r *ghttp.Request, err error) {
	g.Log().Error(err) // 输出log
	if v, ok := err.(gvalid.Error); ok {
		r.Response.WriteJsonExit(model.Response{
			Error: v.Error(),
			Code:  0,
		})
	}
	r.Response.WriteJsonExit(model.Response{
		Error: err.Error(),
		Code:  0,
	})
}
