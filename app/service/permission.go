package service

import (
	"gf-web/app/errcode"
	"gf-web/app/helper"
	"gf-web/app/model/permission_node"
	"gf-web/app/repository"
	"gf-web/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"strings"
	"time"
)

type Permission struct {
	Abstract
	permissionNode   *repository.PermissionNode
	permissionPolicy *repository.PermissionPolicy
}

func NewPermission(request ...*ghttp.Request) *Permission {
	service := &Permission{}

	if len(request) > 0 {
		service.request = request[0]
	}

	service.permissionNode = repository.NewPermissionNode()
	service.permissionPolicy = repository.NewPermissionPolicy()

	return service
}

// 获取权限节点
func (s *Permission) GetNode(id uint) (*permission_node.Entity, error) {
	return s.GetNodeByWhere("id = ?", id)
}

// 拉取节点
func (s *Permission) FetchNode(page int, limit int, where ...interface{}) (*Pagination, error) {
	var (
		entities []*permission_node.Entity
		err      error
		total    int
		fields   = "id,name,path,method,status,remark,created_at,updated_at"
	)

	total, err = s.permissionNode.Count(where...)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	if page <= 0 {
		page = 1
	}

	entities, err = s.permissionNode.Page(page, limit, fields, where...)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	list := make([]g.Map, 0)

	for _, entity := range entities {
		list = append(list, g.Map{
			"id":         entity.Id,
			"name":       entity.Name,
			"path":       entity.Path,
			"method":     entity.Method,
			"status":     entity.Status,
			"remark":     entity.Remark,
			"created_at": entity.CreatedAt,
			"updated_at": entity.UpdatedAt,
		})
	}

	return &Pagination{
		Page:  page,
		Total: total,
		Limit: limit,
		List:  list,
	}, nil
}

// 创建权限节点
func (s *Permission) CreateNode(args g.Map) (*permission_node.Entity, error) {
	exist, err := s.permissionNode.Exist("path = ? and (method = ? or method = ?)", args["path"], args["method"], "ALL")

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	if exist {
		return nil, global.NewError(http.StatusBadRequest, errcode.CodePermissionNodeExist, "节点已存在")
	}

	entity, err := s.permissionNode.Create(args)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	s.SetNodeCache(entity)

	return entity, nil
}

// 更新权限节点
func (s *Permission) UpdateNode(id uint, data g.Map) (*permission_node.Entity, error) {
	var (
		entity *permission_node.Entity
		err    error
	)

	entity, err = s.CheckNodeExistWhenUpdate(id, data)

	if err != nil {
		return nil, err
	}

	if _, err := s.permissionNode.Update(nil, entity, data); err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	if status, ok := data["status"]; ok {
		switch status {
		case permission_node.StatusDisabled:
			s.RemoveNodeCache(entity.Method, entity.Path)
		case permission_node.StatusEnabled:
			s.SetNodeCache(entity)
		}
	}

	return nil, nil
}

// 批量切换权限节点状态
func (s *Permission) BatchSwitchNodeStatus(status uint, ids ...uint) (int64, error) {
	return s.BatchUpdateNode(g.Map{"status": status}, ids...)
}

// 删除权限节点
func (s *Permission) BatchDeleteNode(ids ...uint) (int64, error) {
	var (
		err  error
		rows int64
	)

	_, err = s.DeletePoliciesByNodeIds(ids...)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	rows, err = s.permissionNode.BatchDelete(nil, "id in(?)", ids)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	s.BatchRemoveNodeCache(ids...)

	return rows, nil
}

// 批量更新节点
func (s *Permission) BatchUpdateNode(data g.Map, ids ...uint) (int64, error) {
	var (
		err  error
		rows int64
	)

	if len(ids) > 1 {
		if _, ok := data["method"]; ok {
			return 0, global.NewError(http.StatusBadRequest, errcode.CodeParamsInvalid, "无法批量更新节点操作类型")
		}

		if _, ok := data["path"]; ok {
			return 0, global.NewError(http.StatusBadRequest, errcode.CodeParamsInvalid, "无法批量更新节点路径")
		}
	} else {
		_, err = s.CheckNodeExistWhenUpdate(ids[0], data)

		if err != nil {
			return 0, err
		}
	}

	rows, err = s.permissionNode.BatchUpdate(nil, data, "id in(?)", ids)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	if status, ok := data["status"]; ok {
		switch status {
		case permission_node.StatusDisabled:
			s.BatchRemoveNodeCache(ids...)
		case permission_node.StatusEnabled:
			s.BatchSetNodeCache(ids...)
		}
	}

	return rows, nil
}

// 批量获取节点
func (s *Permission) BatchGetNode(ids ...uint) ([]*permission_node.Entity, error) {
	var (
		err      error
		entities []*permission_node.Entity
	)

	if len(ids) > 0 {
		entities, err = s.permissionNode.BatchGet(nil, "id in(?)", ids)
	} else {
		entities, err = s.permissionNode.BatchGet(nil)
	}

	if err != nil {
		return entities, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	return entities, nil
}

// 根据PATH和METHOD获取节点ID
func (s *Permission) GetNodeIdByUrl(method string, path string) (string, error) {
	var (
		err    error
		entity *permission_node.Entity
	)

	nodeId := s.GetNodeCache(method, path)

	if nodeId != "" {
		return nodeId, nil
	}

	entity, err = s.GetNodeByWhere("method = ? and path = ?", method, path)

	if err != nil {
		return "", err
	}

	s.SetNodeCache(entity)

	return gconv.String(entity.Id), nil
}

// 更新时检测节点是否存在
func (s *Permission) CheckNodeExistWhenUpdate(id uint, data g.Map) (*permission_node.Entity, error) {
	var (
		err    error
		entity *permission_node.Entity
		method string
		path   string
	)

	entity, err = s.GetNode(id)

	if err != nil {
		return nil, err
	}

	if tmpMethod, ok := data["method"]; ok {
		method = gconv.String(tmpMethod)
	} else {
		method = entity.Method
	}

	if tmpPath, ok := data["path"]; ok {
		path = gconv.String(tmpPath)
	} else {
		path = entity.Path
	}

	exist, err := s.permissionNode.Exist("id != ? and path = ? and (method = ? or method = ?)", id, path, method, "ALL")

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	if exist {
		return nil, global.NewError(http.StatusBadRequest, errcode.CodePermissionNodeExist, "节点已存在")
	}

	return entity, nil
}

// 根据条件获取权限节点
func (s *Permission) GetNodeByWhere(where ...interface{}) (*permission_node.Entity, error) {
	var (
		entity *permission_node.Entity
		err    error
	)

	entity, err = s.permissionNode.Get(where...)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	if entity == nil {
		return nil, global.NewError(http.StatusNotFound, errcode.CodeAdminNotExist, "权限节点不存在")
	}

	return entity, nil
}

// 设置节点缓存
func (s *Permission) SetNodeCache(entity *permission_node.Entity) {
	_ = gcache.New().Set(strings.ToLower(entity.Method+entity.Path), entity.Id, time.Hour)
}

// 获取节点缓存
func (s *Permission) GetNodeCache(method string, path string) string {
	if id, err := gcache.New().Get(strings.ToLower(method + path)); err != nil {
		return ""
	} else {
		return gconv.String(id)
	}
}

// 移除节点缓存
func (s *Permission) RemoveNodeCache(method string, path string) {
	_, _ = gcache.New().Remove(strings.ToLower(method + path))
}

// 批量移除节点缓存
func (s *Permission) BatchRemoveNodeCache(ids ...uint) {
	entities, _ := s.BatchGetNode(ids...)

	for _, entity := range entities {
		s.RemoveNodeCache(entity.Method, entity.Path)
	}
}

// 批量设置节点缓存
func (s *Permission) BatchSetNodeCache(ids ...uint) {
	entities, _ := s.BatchGetNode(ids...)

	for _, entity := range entities {
		s.SetNodeCache(entity)
	}
}

func (s *Permission) CreatePolicy() {
	casbin := helper.Casbin(helper.CasbinGroupBackend)

	//casbin.AddGroupingPolicy()
	casbin.RemoveFilteredPolicy(0, "alice")
}

// 创建角色节点策略
func (s *Permission) CreatePolicies(roleId uint, nodeIds ...uint) (bool, error) {
	var (
		err    error
		casbin = helper.Casbin(helper.CasbinGroupBackend)
		rules  = make([][]string, 0)
	)

	for _, nodeId := range gconv.SliceInt(nodeIds) {
		rules = append(rules, []string{
			gconv.String(roleId),
			gconv.String(nodeId),
		})
	}

	if _, err = casbin.AddPolicies(rules); err != nil {
		return false, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	return true, nil
}

// 创建用户角色策略
func (s *Permission) CreateGroupingPolicies(userId uint, roleIds ...uint) (bool, error) {
	var (
		err    error
		casbin = helper.Casbin(helper.CasbinGroupBackend)
		rules  = make([][]string, 0)
	)

	for _, roleId := range gconv.SliceInt(roleIds) {
		rules = append(rules, []string{
			gconv.String(userId),
			gconv.String(roleId),
		})
	}

	if _, err = casbin.AddGroupingPolicies(rules); err != nil {
		return false, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	return true, nil
}

// 根据角色ID删除角色节点策略
func (s *Permission) DeletePoliciesByRoleIds(roleIds ...uint) (bool, error) {
	return s.DeletePolicies(0, roleIds...)
}

// 根据节点ID删除角色节点策略
func (s *Permission) DeletePoliciesByNodeIds(nodeIds ...uint) (bool, error) {
	return s.DeletePolicies(1, nodeIds...)
}

// 根据角色ID获取角色节点策略
func (s *Permission) GetPoliciesByRoleIds(roleIds ...uint) [][]string {
	return s.GetPolicies(0, roleIds...)
}

// 根据节点ID获取角色节点策略
func (s *Permission) GetPoliciesByNodeIds(nodeIds ...uint) [][]string {
	return s.GetPolicies(1, nodeIds...)
}

// 根据用户ID删除用户角色策略组
func (s *Permission) DeleteGroupingPoliciesByUserIds(userIds ...uint) (bool, error) {
	return s.DeleteGroupingPolicies(0, userIds...)
}

// 根据角色ID删除用户角色策略组
func (s *Permission) DeleteGroupingPoliciesByRoleIds(roleIds ...uint) (bool, error) {
	return s.DeleteGroupingPolicies(1, roleIds...)
}

// 根据用户ID获取用户角色策略组
func (s *Permission) GetGroupingPoliciesByUserIds(userIds ...uint) [][]string {
	return s.GetGroupingPolicies(0, userIds...)
}

// 根据角色ID获取用户角色策略组
func (s *Permission) GetGroupingPoliciesByRoleIds(roleIds ...uint) [][]string {
	return s.GetGroupingPolicies(1, roleIds...)
}

// 删除策略
func (s *Permission) DeletePolicies(fieldIndex int, ids ...uint) (bool, error) {
	var (
		err    error
		casbin = helper.Casbin(helper.CasbinGroupBackend)
		rules  = s.GetPolicies(fieldIndex, ids...)
	)

	_, err = casbin.RemovePolicies(rules)

	if err != nil {
		return false, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	return true, nil
}

// 删除策略组
func (s *Permission) DeleteGroupingPolicies(fieldIndex int, ids ...uint) (bool, error) {
	var (
		err    error
		casbin = helper.Casbin(helper.CasbinGroupBackend)
		rules  = s.GetGroupingPolicies(fieldIndex, ids...)
	)

	_, err = casbin.RemoveGroupingPolicies(rules)

	if err != nil {
		return false, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	return true, nil
}

// 获取策略
func (s *Permission) GetPolicies(fieldIndex int, ids ...uint) [][]string {
	var (
		casbin = helper.Casbin(helper.CasbinGroupBackend)
		rules  = make([][]string, 0)
	)

	for _, id := range ids {
		for _, item := range casbin.GetFilteredNamedPolicy("p", fieldIndex, gconv.String(id)) {
			rules = append(rules, item)
		}
	}

	return rules
}

// 获取策略组
func (s *Permission) GetGroupingPolicies(fieldIndex int, ids ...uint) [][]string {
	var (
		casbin = helper.Casbin(helper.CasbinGroupBackend)
		rules  = make([][]string, 0)
	)

	for _, id := range ids {
		for _, item := range casbin.GetFilteredNamedGroupingPolicy("g", fieldIndex, gconv.String(id)) {
			rules = append(rules, item)
		}
	}

	return rules
}
