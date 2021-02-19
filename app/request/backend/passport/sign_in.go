package req_passport

import (
	"gf-web/library/global"
)

type SignIn struct {
	Account  string `c:"account"`
	Email    string `c:"email"`
	Mobile   string `c:"mobile"`
	Password string `c:"password"`
}

func (r *SignIn) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "account",
			Rule:    "required-without:email",
			Message: "账号或邮箱必需",
		},
		{
			Field:   "account",
			Rule:    "min-length:4",
			Message: "账号至少:min个字符",
		},
		{
			Field:   "email",
			Rule:    "required-without:account",
			Message: "账号或邮箱必需",
		},
		{
			Field:   "email",
			Rule:    "email",
			Message: "邮箱不合法",
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
