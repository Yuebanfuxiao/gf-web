package repository

import (
	"database/sql"
	"gf-web/app/model/admin_login_record"
	"github.com/gogf/gf/util/gconv"
)

type AdminLoginRecord struct {
}

/**
 * 创建数据仓库
 */
func NewAdminLoginRecord() *AdminLoginRecord {
	repository := &AdminLoginRecord{}

	return repository
}

/**
 * 创建数据
 */
func (r *AdminLoginRecord) Create(data ...interface{}) (sql.Result, error) {
	return admin_login_record.Model.Insert(data)
}

// 插入数据
func (r *AdminLoginRecord) Insert(entity *admin_login_record.Entity) (*admin_login_record.Entity, error) {
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

/**
 * 保存数据
 */
func (r *AdminLoginRecord) Save(entity *admin_login_record.Entity) (sql.Result, error) {
	return entity.Save()
}
