package middleware

import (
	"gf-web/app/errcode"
	"gf-web/app/helper"
	"gf-web/library/global"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

// JWT授权验证(后端)
func JwtAuthBackend(request *ghttp.Request) {
	jwtAuth(request, helper.Jwt(helper.JwtGroupBackend))
}

// JWT授权验证(前端)
func JwtAuthFrontend(request *ghttp.Request) {
	jwtAuth(request, helper.Jwt(helper.JwtGroupFrontend))
}

/**
 * JWT授权验证
 */
func jwtAuth(request *ghttp.Request, jwt *global.Jwt) {
	if err := jwt.Middleware(request); err != nil {
		response := helper.Response(request).WithHeader("WWW-Authenticate", "JWT realm="+jwt.Realm)

		if e, ok := err.(*global.JwtError); ok {
			switch e.Errors {
			case global.JwtErrorExpiredToken:
				response.Response(http.StatusUnauthorized, errcode.CodeTokenExpired, "授权已过期")
			case global.JwtErrorMissingExpField, global.JwtErrorWrongFormatOfExp:
				response.Response(http.StatusUnauthorized, errcode.CodeTokenInvalid, "无效授权")
			case global.JwtErrorAuthorizeElsewhere:
				response.Response(http.StatusUnauthorized, errcode.CodeAuthorizeElsewhere, "账号在其它地方登陆")
			}
		}

		response.Response(http.StatusUnauthorized, errcode.CodeUnauthorized, "未授权")
	}

	request.Middleware.Next()
}
