package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

// 跨域中间件
func Cors(request *ghttp.Request) {
	corsOptions := request.Response.DefaultCORSOptions()

	corsOptions.AllowDomain = []string{"goframe.org", "baidu.com"}

	if !request.Response.CORSAllowedOrigin(corsOptions) {
		request.Response.WriteStatus(http.StatusForbidden)
		return
	}

	request.Response.CORS(corsOptions)

	request.Middleware.Next()
}
