package service

import (
	"gf-web/app/errcode"
	"gf-web/app/model/role"
	"gf-web/app/repository"
	"gf-web/library/global"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"net/http"
)

type Role struct {
	Abstract
	role *repository.Role
}

type UpdateRoleArgs struct {
	Name        string
	Description string
	Status      uint
}

func NewRole(request ...*ghttp.Request) *Role {
	service := &Role{}

	if request != nil {
		service.request = request[0]
	}

	service.role = repository.NewRole()

	return service
}

// 拉取角色
func (s *Role) FetchRole(args g.Map) (*repository.Pagination, error) {
	var (
		err        error
		pagination *repository.Pagination
		fields     = struct {
			Id        uint
			Name      string
			Remark    string
			Status    uint
			CreatedAt *gtime.Time
			UpdatedAt *gtime.Time
			DeletedAt *gtime.Time
		}{}
	)

	pagination, err = s.role.Page(args, fields)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	return pagination, nil
}

// 创建角色
func (s *Role) CreateRole(args g.Map) (*role.Entity, error) {
	var (
		tx         *gdb.TX
		err        error
		entity     *role.Entity
		permission = NewPermission(s.request)
	)

	if tx, err = g.DB().Begin(); err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}
	}()

	entity, err = s.role.Create(tx, args)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	if nodes, ok := args["nodes"]; ok {
		_, err = permission.CreatePolicies(entity.Id, gconv.SliceUint(nodes)...)

		if err != nil {
			return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
		}
	}

	return entity, nil
}

// 删除角色
func (s *Role) DeleteRole(ids ...uint) (int64, error) {
	var (
		err        error
		rows       int64
		permission = NewPermission(s.request)
	)

	_, err = permission.DeletePoliciesByRoleIds(ids...)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	_, err = permission.DeleteGroupingPoliciesByRoleIds(ids...)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	rows, err = s.role.BatchDelete(nil, "id in(?)", ids)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	return rows, nil
}

// 更新角色
func (s *Role) UpdateRole(id uint, args g.Map) (*role.Entity, error) {
	var (
		err        error
		entity     *role.Entity
		permission = NewPermission(s.request)
	)

	entity, err = s.GetById(id)

	if err != nil {
		return nil, err
	}

	if nodes, ok := args["nodes"]; ok {
		_, err = permission.DeletePoliciesByRoleIds(id)

		if err != nil {
			return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
		}

		_, err = permission.CreatePolicies(id, gconv.SliceUint(nodes)...)

		if err != nil {
			return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
		}
	}

	entity, err = s.role.Update(nil, entity, args)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	return entity, nil
}

// 切换状态
func (s *Role) SwitchStatus(id uint, status uint) (*role.Entity, error) {
	return s.UpdateRole(id, g.Map{"status": status})
}

// 批量获取角色
func (s *Role) BatchGetRole(ids ...uint) ([]*role.Entity, error) {
	var entities []*role.Entity

	entities, err := s.role.BatchGet(nil, "id in(?)", ids)

	if err != nil {
		return entities, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	return entities, nil
}

// 根据ID获取角色
func (s *Role) GetById(id uint) (*role.Entity, error) {
	var (
		entity *role.Entity
		err    error
	)

	entity, err = s.role.Get(id)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	if entity == nil {
		return nil, global.NewError(http.StatusNotFound, errcode.CodeAdminNotExist, "角色不存在")
	}

	return entity, nil
}

// 根据管理员ID获取角色
func (s *Role) GetRolesByAdminId(adminId uint) ([]*role.Entity, error) {
	var (
		permission = NewPermission(s.request)
		ids        = make([]uint, 0)
	)

	rules := permission.GetGroupingPoliciesByUserIds(adminId)

	for _, rule := range rules {
		ids = append(ids, gconv.Uint(rule[1]))
	}

	return s.BatchGetRole(ids...)
}
