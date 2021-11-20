package middleware

import "github.com/gogf/gf/net/ghttp"

// 允许跨域请求中间件
func MiddlewareCors(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = []string{"goframe.org", "johng.cn", "localhost"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}
