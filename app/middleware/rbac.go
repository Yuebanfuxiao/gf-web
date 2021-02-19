package middleware

import (
	"gf-web/app/errcode"
	"gf-web/app/helper"
	"gf-web/app/service"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"net/http"
)

// RBAC权限验证(后端)
func RbacBackend(request *ghttp.Request) {
	var (
		hasPermission = false
		jwt           = helper.Jwt(helper.JwtGroupBackend)
		res           = helper.Response(request)
		casbin        = helper.Casbin(helper.CasbinGroupBackend)
		adminId       = jwt.GetIdentity(request)
		permission    = service.NewPermission(request)
		nodeId, err   = permission.GetNodeIdByUrl(request.Method, request.URL.Path)
	)

	if adminId == nil {
		res.Response(http.StatusForbidden, errcode.CodeNoPermission, "权限不足")
	}

	if err != nil {
		res.Response(http.StatusForbidden, errcode.CodeNoPermission, "权限不足")
	}

	if b, err := casbin.Enforce(gconv.String(adminId), nodeId); err == nil && b {
		hasPermission = true
	}

	if !hasPermission {
		res.Response(http.StatusForbidden, errcode.CodeNoPermission, "权限不足")
	}

	request.Middleware.Next()
}
