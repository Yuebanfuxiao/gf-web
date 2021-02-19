package req_admin

import (
	"gf-web/library/global"
)

type CreateAdmin struct {
	Account  string `c:"account"`
	Email    string `c:"email"`
	Mobile   string `c:"mobile"`
	Password string `c:"password"`
	Nickname string `c:"nickname"`
	Avatar   string `c:"avatar"`
	Status   uint   `c:"status"`
}

func (r *CreateAdmin) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "account",
			Rule:    "required-without-all:email,mobile",
			Message: "账号必需",
		},
		{
			Field:   "account",
			Rule:    "min-length:4",
			Message: "账号至少:min个字符",
		},
		{
			Field:   "email",
			Rule:    "required-without-all:account,mobile",
			Message: "邮箱必需",
		},
		{
			Field:   "email",
			Rule:    "email",
			Message: "邮箱不合法",
		},
		{
			Field:   "mobile",
			Rule:    "required-without-all:account,email",
			Message: "手机号必需",
		},
		{
			Field:   "mobile",
			Rule:    "phone",
			Message: "手机号不合法",
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
		{
			Field:   "nickname",
			Rule:    "required",
			Message: "昵称必需",
		},
		{
			Field:   "nickname",
			Rule:    "max-length:20",
			Message: "昵称至多:max个字符",
		},
		{
			Field:   "avatar",
			Rule:    "required",
			Message: "头像必需",
		},
		{
			Field:   "avatar",
			Rule:    "max-length:200",
			Message: "头像不合法",
		},
		{
			Field:   "status",
			Rule:    "required",
			Message: "管理员状态必需",
		},
		{
			Field:   "status",
			Rule:    "in:0,1",
			Message: "管理员状态非法",
		},
	}
}
