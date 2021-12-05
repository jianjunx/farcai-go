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
		// 添加文章
		group.POST("/post", api.Post.AddPost)
		// 更新文章
		group.PUT("/post", api.Post.AddPost)
		// 根据ID查询文章
		group.GET("/post/:id", api.Post.GetPost)
		// 删除文章
		group.DELETE("/post/:id", api.Post.DeletePost)
		// 上传文件
		group.GET("/credentials/cos", api.Assets.COSCredentials)
	})
	// html
	s.BindHandler("/", api.Html.Home)
	s.BindHandler("/p/:id", api.Html.Detail)
	s.BindHandler("/writing", api.Html.Writing)
	s.BindHandler("/login", api.Html.Login)
}
