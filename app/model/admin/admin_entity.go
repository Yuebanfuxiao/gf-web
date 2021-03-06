// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package admin

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table admin.
type Entity struct {
    Id          uint        `orm:"id,primary"    json:"id"`            // 邮箱                
    Account     string      `orm:"account"       json:"account"`       // 账号                
    Mobile      string      `orm:"mobile"        json:"mobile"`        // 手机号              
    Email       string      `orm:"email"         json:"email"`         // 邮箱                
    Password    string      `orm:"password"      json:"password"`      // 密码                
    Nickname    string      `orm:"nickname"      json:"nickname"`      // 昵称                
    Avatar      string      `orm:"avatar"        json:"avatar"`        // 头像地址            
    RegisterAt  *gtime.Time `orm:"register_at"   json:"register_at"`   // 注册时间            
    RegisterIp  string      `orm:"register_ip"   json:"register_ip"`   // 注册IP              
    LastLoginAt *gtime.Time `orm:"last_login_at" json:"last_login_at"` // 最后登陆时间        
    LastLoginIp string      `orm:"last_login_ip" json:"last_login_ip"` // 最后登陆IP          
    Status      uint        `orm:"status"        json:"status"`        // 状态 1:启用 0:禁用  
    CreatedAt   *gtime.Time `orm:"created_at"    json:"created_at"`    //                     
    UpdatedAt   *gtime.Time `orm:"updated_at"    json:"updated_at"`    //                     
    DeletedAt   *gtime.Time `orm:"deleted_at"    json:"deleted_at"`    //                     
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}