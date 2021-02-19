package req_admin

import (
	"gf-web/library/global"
)

type ResetPassword struct {
	Id       uint   `c:"id"`
	Password string `c:"password"`
}

func (r *ResetPassword) Rules() global.Rules {
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
			Field:   "password",
			Rule:    "required",
			Message: "密码必需",
		},
		{
			Field:   "password",
			Rule:    "min-length:6",
			Message: "密码至少:min个字符",
		},
		{
			Field:   "password",
			Rule:    "max-length:22",
			Message: "密码至多:max个字符",
		},
	}
}
