package backend

import (
	"gf-web/app/helper"
	"gf-web/app/request/backend/admin"
	"gf-web/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type Admin struct {
}

func (c *Admin) GetList(request *ghttp.Request) {

}

// 创建管理员
func (c *Admin) CreateAdmin(request *ghttp.Request) {
	var (
		req  req_admin.CreateAdmin
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if _, err := service.NewAdmin(request).CreateAdmin(args); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "创建成功")
	}
}

// 删除管理员
func (c *Admin) DeleteAdmin(request *ghttp.Request) {
	var (
		req req_admin.DeleteAdmin
	)

	helper.Accept(request, &req)

	if _, err := service.NewAdmin(request).BatchDeleteAdmin(req.Id...); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "删除成功")
	}
}

// 修改管理员
func (c *Admin) UpdateAdmin(request *ghttp.Request) {
	var (
		req  req_admin.UpdateAdmin
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if _, err := service.NewAdmin(request).UpdateAdmin(req.Id, args); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "更新成功")
	}
}

// 切换状态
func (c *Admin) SwitchStatus(request *ghttp.Request) {
	var (
		req req_admin.SwitchStatus
	)

	helper.Accept(request, &req)

	if _, err := service.NewAdmin(request).BatchSwitchStatus(req.Status, req.Id...); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "切换成功")
	}
}

// 重置密码
func (c *Admin) ResetPassword(request *ghttp.Request) {
	var (
		req req_admin.ResetPassword
	)

	helper.Accept(request, &req)

	if _, err := service.NewAdmin(request).ResetPassword(req.Id, req.Password); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "重置成功")
	}
}

// 修改密码(当前管理员)
func (c *Admin) UpdatePassword(request *ghttp.Request) {
	var (
		req req_admin.UpdatePassword
	)

	helper.Accept(request, &req)

	id := helper.Jwt(helper.JwtGroupBackend).GetIdentity(request)

	if _, err := service.NewAdmin(request).UpdatePassword(gconv.Uint(id), req.OldPassword, req.NewPassword); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "修改成功")
	}
}
