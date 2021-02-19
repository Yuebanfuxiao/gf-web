package backend

import (
	"gf-web/app/helper"
	"gf-web/app/request/backend/role"
	"gf-web/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Role struct {
}

// 拉取角色
func (c *Role) FetchRole(request *ghttp.Request) {
	var (
		req  req_role.FetchRole
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if ret, err := service.NewRole(request).FetchRole(args); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "获取成功", ret)
	}
}

// 创建角色
func (c *Role) CreateRole(request *ghttp.Request) {
	var (
		req  req_role.CreateRole
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if _, err := service.NewRole(request).CreateRole(args); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "创建成功")
	}
}

// 删除角色
func (c *Role) DeleteRole(request *ghttp.Request) {
	var (
		req req_role.DeleteRole
	)

	helper.Accept(request, &req)

	if _, err := service.NewRole(request).DeleteRole(req.Id...); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "删除成功")
	}
}

// 更新角色
func (c *Role) UpdateRole(request *ghttp.Request) {
	var (
		req  req_role.UpdateRole
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if _, err := service.NewRole(request).UpdateRole(req.Id, args); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "更新成功")
	}
}

// 切换状态
func (c *Role) SwitchStatus(request *ghttp.Request) {
	var (
		req req_role.SwitchStatus
	)

	helper.Accept(request, &req)

	if _, err := service.NewRole(request).SwitchStatus(req.Id, req.Status); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "切换成功")
	}
}
