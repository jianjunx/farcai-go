package middleware

import (
	"errors"
	"farcai-go/app/service"
	"farcai-go/library/jwt"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareAuth(r *ghttp.Request) {
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		r.Response.WriteStatus(http.StatusForbidden)
		service.ErrorHandler(r, errors.New("token为空"))
		return
	}
	token, claims, err := jwt.ParseToken(authToken)
	if err != nil || !token.Valid {
		r.Response.WriteStatus(http.StatusForbidden)
		service.ErrorHandler(r, errors.New("身份认证失败"))
		return
	}
	r.SetCtxVar("UserName", claims.UserName)
	r.Middleware.Next()
}
