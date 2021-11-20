package router

import (
	"farcai-go/app/api"
	"farcai-go/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.MiddlewareCors)
		// 注册
		group.POST("/signup", api.User.Register)
		// 登录
		group.POST("/login", api.User.Login)
		// 认证中间件
		group.Middleware(middleware.MiddlewareAuth)
		// 获取用户信息
		group.GET("/user-info", api.User.GetUserInfo)
		// 添加分类
		group.POST("/category", api.Category.AddCategory)
		// 获取分类列表
		group.GET("/category", api.Category.GetCategorys)

	})
}
