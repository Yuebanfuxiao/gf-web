package backend

import (
	"gf-web/app/helper"
	"gf-web/app/request/backend/permission"
	"gf-web/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Permission struct {
}

// 列表节点
func (c *Permission) FetchNode(request *ghttp.Request) {
	var (
		req req_permission.FetchNode
	)

	helper.Accept(request, &req)

	if ret, err := service.NewPermission(request).FetchNode(req.Page, req.Limit); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "获取成功", ret)
	}
}

// 创建节点
func (c *Permission) CreateNode(request *ghttp.Request) {
	var (
		req  req_permission.CreateNode
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if _, err := service.NewPermission(request).CreateNode(args); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "创建成功")
	}
}

// 删除节点（批量）
func (c *Permission) DeleteNode(request *ghttp.Request) {
	var (
		req req_permission.DeleteNode
	)

	helper.Accept(request, &req)

	if _, err := service.NewPermission(request).BatchDeleteNode(req.Id...); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "删除成功")
	}
}

// 更新节点
func (c *Permission) UpdateNode(request *ghttp.Request) {
	var (
		req  req_permission.UpdateNode
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if _, err := service.NewPermission(request).UpdateNode(req.Id, args); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "更新成功")
	}
}

// 切换节点状态（批量）
func (c *Permission) SwitchNodeStatus(request *ghttp.Request) {
	var (
		req req_permission.SwitchNodeStatus
	)

	helper.Accept(request, &req)

	if _, err := service.NewPermission(request).BatchSwitchNodeStatus(req.Status, req.Id...); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "切换成功")
	}
}
