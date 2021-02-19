package repository

import (
	"database/sql"
	"gf-web/app/model/admin"
	"gf-web/app/model/permission_node"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type PermissionNode struct {
	Abstract
}

// 创建数据仓库
func NewPermissionNode() *PermissionNode {
	repository := &PermissionNode{}

	return repository
}

// 创建数据
func (r *PermissionNode) Create(data interface{}) (*permission_node.Entity, error) {
	var (
		entity permission_node.Entity
		result sql.Result
		err    error
	)

	if result, err = permission_node.Model.Data(data).Filter().Insert(); err != nil {
		return nil, err
	} else {
		if id, err := result.LastInsertId(); err != nil {
			return nil, err
		} else {
			_ = gconv.Struct(data, &entity)

			entity.Id = gconv.Uint(id)

			return &entity, nil
		}
	}
}

// 更新权限节点
func (r *PermissionNode) Update(tx *gdb.TX, entity *permission_node.Entity, data g.Map) (*permission_node.Entity, error) {
	args := gconv.MapDeep(data)

	_, err := r.BatchUpdate(tx, args, "id = ?", entity.Id)

	if err != nil {
		return nil, err
	}

	_ = gconv.Struct(args, entity)

	return entity, nil
}

// 删除权限节点
func (r *PermissionNode) Delete(entity *admin.Entity) (int64, error) {
	result, err := entity.Delete()

	if err != nil {
		return 0, err
	}

	return r.handleDeleteResult(result)
}

// 分页数据
func (r *PermissionNode) Page(page int, limit int, fields interface{}, where ...interface{}) ([]*permission_node.Entity, error) {
	return permission_node.Model.Fields(fields).Page(page, limit).All(where...)
}

// 检测权限节点是否存在
func (r *PermissionNode) Exist(where ...interface{}) (bool, error) {
	if count, err := r.Count(where...); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

// 获取权限节点数量
func (r *PermissionNode) Count(where ...interface{}) (int, error) {
	return permission_node.Model.Count(where...)
}

// 获取权限节点
func (r *PermissionNode) Get(where ...interface{}) (*permission_node.Entity, error) {
	return permission_node.Model.One(where)
}

// 批量删除权限节点
func (r *PermissionNode) BatchDelete(tx *gdb.TX, where ...interface{}) (int64, error) {
	model := permission_node.Model

	if tx != nil {
		model = model.TX(tx)
	}

	result, err := model.Delete(where...)

	if err != nil {
		return 0, err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return 0, err
	} else {
		return rows, nil
	}
}

// 批量更新权限节点
func (r *PermissionNode) BatchUpdate(tx *gdb.TX, dataAndWhere ...interface{}) (int64, error) {
	model := admin.Model

	if tx != nil {
		model = model.TX(tx)
	}

	result, err := model.Filter().Update(dataAndWhere...)

	if err != nil {
		return 0, err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return 0, err
	} else {
		return rows, nil
	}
}

// 批量获取权限
func (r *PermissionNode) BatchGet(tx *gdb.TX, where ...interface{}) ([]*permission_node.Entity, error) {
	model := permission_node.Model

	if tx != nil {
		model = model.TX(tx)
	}

	return model.All(where...)
}
