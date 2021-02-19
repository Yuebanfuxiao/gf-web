package helper

import (
	"gf-web/app/component"
	"gf-web/library/global"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/net/ghttp"
)

const (
	JwtGroupBackend    = "backend"
	JwtGroupFrontend   = "frontend"
	CasbinGroupBackend = "backend"
)

// Request returns an instance of request with specified configuration group name.
func Request(request *ghttp.Request) *global.Request {
	return global.NewRequest(request)
}

// Response returns an instance of response with specified configuration group name.
func Response(name ...interface{}) *global.Response {
	return component.Response(name...)
}

// JWT returns an instance of jwt with specified configuration group name.
func Jwt(name ...string) *global.Jwt {
	return component.Jwt(name...)
}

// Casbin returns an instance of casbin with specified configuration group name.
func Casbin(name ...string) *casbin.Enforcer {
	return component.Casbin(name...)
}