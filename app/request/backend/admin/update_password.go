package req_admin

import (
	"gf-web/library/global"
)

type UpdatePassword struct {
	OldPassword string `c:"old_password"`
	NewPassword string `c:"new_password"`
}

func (r *UpdatePassword) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "old_password",
			Rule:    "required",
			Message: "旧密码必需",
		},
		{
			Field:   "old_password",
			Rule:    "min-length:6",
			Message: "旧密码至少:min个字符",
		},
		{
			Field:   "old_password",
			Rule:    "max-length:22",
			Message: "旧密码至多:max个字符",
		},
		{
			Field:   "new_password",
			Rule:    "required",
			Message: "新密码必需",
		},
		{
			Field:   "new_password",
			Rule:    "min-length:6",
			Message: "新密码至少:min个字符",
		},
		{
			Field:   "new_password",
			Rule:    "max-length:22",
			Message: "新密码至多:max个字符",
		},
	}
}
