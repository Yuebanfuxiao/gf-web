package repository

import (
	"database/sql"
	"fmt"
	"gf-web/app/model/role"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

type Role struct {
	Abstract
}

// 创建数据仓库
func NewRole() *Role {
	repository := &Role{}

	return repository
}

// 根据ID获取角色
func (r *Role) Get(id uint) (*role.Entity, error) {
	return role.Model.Where("id", id).One()
}

// 创建数据
func (r *Role) Create(tx *gdb.TX, data interface{}) (*role.Entity, error) {
	var (
		entity role.Entity
		result sql.Result
		err    error
		args   = gconv.MapDeep(data)
	)

	model := role.Model

	if tx != nil {
		model = model.TX(tx)
	}

	if result, err = model.Data(args).Filter().Insert(); err != nil {
		return nil, err
	} else {
		if id, err := result.LastInsertId(); err != nil {
			return nil, err
		} else {
			_ = gconv.Struct(args, &entity)

			entity.Id = gconv.Uint(id)

			return &entity, nil
		}
	}
}

// 删除角色
func (r *Role) Delete(tx *gdb.TX, entity *role.Entity) (bool, error) {
	_, err := r.BatchDelete(tx, "id = ?", entity.Id)

	if err != nil {
		return false, err
	}

	return true, nil
}

// 更新角色
func (r *Role) Update(tx *gdb.TX, entity *role.Entity, data g.Map) (*role.Entity, error) {
	args := gconv.MapDeep(data)

	_, err := r.BatchUpdate(tx, args, "id = ?", entity.Id)

	if err != nil {
		return nil, err
	}

	_ = gconv.Struct(args, entity)

	return entity, nil
}

// 插入数据
func (r *Role) Insert(entity *role.Entity) (*role.Entity, error) {
	if result, err := entity.Insert(); err != nil {
		return nil, err
	} else {
		if id, err := result.LastInsertId(); err != nil {
			return nil, err
		} else {
			entity.Id = gconv.Uint(id)
		}

		return entity, nil
	}
}

// 更新或插入角色
func (r *Role) Save(entity *role.Entity) (*role.Entity, error) {
	result, err := entity.Save()

	if err != nil {
		return nil, err
	}

	if _, id, err := r.handleSaveResult(result); err != nil {
		return nil, err
	} else {
		if id != nil {
			entity.Id = id.(uint)
		}

		return entity, nil
	}
}

// 角色分页数据
func (r *Role) Page(args g.Map, fields interface{}, where ...interface{}) (*Pagination, error) {
	var (
		err      error
		page     int
		limit    int
		total    int
		entities []*role.Entity
		list     []g.Map
		filter   = role.NewFilter(args)
	)

	total, err = filter.Model.Count()

	if err != nil {
		return nil, err
	}

	entities, err = filter.Model.Fields(fields).All(where...)

	fmt.Println(entities)

	list, err = r.handleList(entities, fields)

	return &Pagination{
		Page:  page,
		Total: total,
		Limit: limit,
		List:  list,
	}, nil
}

// 获取角色数量
func (r *Role) Count(where ...interface{}) (int, error) {
	return role.Model.Count(where...)
}

// 批量获取角色
func (r *Role) BatchGet(tx *gdb.TX, where ...interface{}) ([]*role.Entity, error) {
	model := role.Model

	if tx != nil {
		model = model.TX(tx)
	}

	return model.All(where...)
}

// 批量创建角色
func (r *Role) BatchCreate(tx *gdb.TX, list g.List) (int64, error) {
	model := role.Model

	if tx != nil {
		model = model.TX(tx)
	}

	result, err := model.Data(list).Filter().Insert()

	if err != nil {
		return 0, err
	}

	if rows, err := result.RowsAffected(); err != nil {
		return 0, err
	} else {
		return rows, nil
	}
}

// 批量删除角色
func (r *Role) BatchDelete(tx *gdb.TX, where ...interface{}) (int64, error) {
	model := role.Model

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

// 批量更新角色
func (r *Role) BatchUpdate(tx *gdb.TX, dataAndWhere ...interface{}) (int64, error) {
	model := role.Model

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
