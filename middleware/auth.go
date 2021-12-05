package middleware

import (
	"farcai-go/app/model"
	"farcai-go/library/jwt"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareAuth(r *ghttp.Request) {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		r.Response.WriteJsonExit(model.Response{
			Error: "token为空",
			Code:  403,
		})
		g.Log().Print("token为空")
		return
	}

	token, claims, err := jwt.ParseToken(authToken)
	// || r.GetClientIp() != claims.Audience
	if err != nil || !token.Valid {
		r.Response.WriteJsonExit(model.Response{
			Error: "身份认证失败",
			Code:  403,
		})
		g.Log().Fatalf("身份认证失败")
		return
	}
	r.SetCtxVar("uid", claims.Uid)
	r.Middleware.Next()
}
