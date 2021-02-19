package helper

import (
	"gf-web/app/errcode"
	"gf-web/library/global"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

// 接收参数
func Accept(request *ghttp.Request, req global.IRequest, args ...interface{}) {
	if err := Request(request).Accept(req, args...); err != nil {
		Response(request).Response(http.StatusUnprocessableEntity, errcode.CodeParameterVerificationFailed, err.Error())
	}
}

// 失败响应
func Fail(request *ghttp.Request, err error) {
	if e, ok := err.(*global.Error); ok {
		Response(request).Response(e.Status(), e.Code(), e.Error())
	} else {
		Response(request).Response(http.StatusInternalServerError, errcode.CodeServerException, err.Error())
	}
}

// 成功响应
func Success(request *ghttp.Request, message string, data ...interface{}) {
	Response(request).Success(message, data...)
}