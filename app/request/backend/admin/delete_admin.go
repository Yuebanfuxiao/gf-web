package req_admin

import (
	"gf-web/library/global"
)

type DeleteAdmin struct {
	Id []uint `c:"id"`
}

func (r *DeleteAdmin) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "id",
			Rule:    "required",
			Message: "管理员ID必需",
		},
		{
			Field:   "id",
			Rule:    "array:integer,1,-1,1",
			Message: "管理员ID非法",
		},
	}
}
