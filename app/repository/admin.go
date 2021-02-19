package repository

import (
	"database/sql"
	"fmt"
	"gf-web/app/model/admin"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

type Admin struct {
	Abstract
}

// 创建数据仓库
func NewAdmin(request ...*ghttp.Request) *Admin {
	repository := &Admin{}

	if len(request) > 0 {
		repository.WithRequest(request[0])
	}

	return repository
}

// 根据ID获取管理员
func (r *Admin) GetById(id uint) (*admin.Entity, error) {
	return admin.Model.Where("id", id).One()
}

// 根据账号获取管理员
func (r *Admin) GetByAccount(account string) (*admin.Entity, error) {
	return admin.Model.Where("account", account).One()
}

// 根据邮箱获取管理员
func (r *Admin) GetByEmail(email string) (*admin.Entity, error) {
	return admin.Model.Where("email", email).One()
}

// 根据手机号获取管理员
func (r *Admin) GetByMobile(mobile string) (*admin.Entity, error) {
	return admin.Model.Where("mobile", mobile).One()
}

// 获取单个管理员
func (r *Admin) Get(where ...interface{}) (*admin.Entity, error) {
	return admin.Model.One(where...)
}

// 获取管理员数量
func (r *Admin) Count(where ...interface{}) (int, error) {
	return admin.Model.Count(where...)
}

// 检测管理员是否存在
func (r *Admin) Exist(where ...interface{}) (bool, error) {
	if count, err := r.Count(where...); err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

// 更新或插入数据
func (r *Admin) Save(entity *admin.Entity) (*admin.Entity, error) {
	data := gconv.Map(entity)

	fmt.Println("_____________________")
	fmt.Println(data)
	fmt.Println("_____________________")

	result, err := admin.Model.Data(data).Filter().Save()

	if err != nil {
		return nil, err
	}

	fmt.Println("_____________________")
	fmt.Println(data)
	fmt.Println("_____________________")

	if _, id, err := r.handleSaveResult(result); err != nil {
		return nil, err
	} else {
		_ = gconv.Struct(data, entity)

		if id != nil {
			entity.Id = id.(uint)
		}

		return entity, nil
	}
}

// 创建数据
func (r *Admin) Create(tx *gdb.TX, data interface{}) (*admin.Entity, error) {
	var (
		entity admin.Entity
		result sql.Result
		err    error
		args   = admin.Creating(gconv.MapDeep(data), r.request)
	)

	model := admin.Model

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

// 更新管理员
func (r *Admin) Update(tx *gdb.TX, entity *admin.Entity, data g.Map) (*admin.Entity, error) {
	args := gconv.MapDeep(data)

	_, err := r.BatchUpdate(tx, args, "id = ?", entity.Id)

	if err != nil {
		return nil, err
	}

	_ = gconv.Struct(args, entity)

	return entity, nil
}

// 删除管理员
func (r *Admin) Delete(entity *admin.Entity) (int64, error) {
	result, err := entity.Delete()

	if err != nil {
		return 0, err
	}

	return r.handleDeleteResult(result)
}

// 批量创建管理员
func (r *Admin) BatchCreate(tx *gdb.TX, list g.List) (int64, error) {
	model := admin.Model

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

// 批量删除管理员
func (r *Admin) BatchDelete(tx *gdb.TX, where ...interface{}) (int64, error) {
	model := admin.Model

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

// 批量更新管理员
func (r *Admin) BatchUpdate(tx *gdb.TX, dataAndWhere ...interface{}) (int64, error) {
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
