package req_admin

import (
	"gf-web/library/global"
)

type UpdateAdmin struct {
	Id       uint   `c:"id"`
	Nickname string `c:"nickname"`
	Avatar   string `c:"avatar"`
	Status   uint   `c:"status"`
}

func (r *UpdateAdmin) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "id",
			Rule:    "required",
			Message: "管理员ID必需",
		},
		{
			Field:   "id",
			Rule:    "integer",
			Message: "管理员ID非法",
		},
		{
			Field:   "id",
			Rule:    "min:1",
			Message: "管理员ID非法",
		},
		{
			Field:   "nickname",
			Rule:    "max-length:20",
			Message: "昵称至多:max个字符",
		},
		{
			Field:   "avatar",
			Rule:    "max-length:200",
			Message: "头像不合法",
		},
		{
			Field:   "status",
			Rule:    "in:0,1",
			Message: "管理员状态非法",
		},
	}
}
