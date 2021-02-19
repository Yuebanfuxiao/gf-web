package req_role

import (
	"gf-web/library/global"
)

type DeleteRole struct {
	Id []uint
}

func (r *DeleteRole) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "id",
			Rule:    "required",
			Message: "角色ID必需",
		},
		{
			Field:   "id",
			Rule:    "array:integer,1,-1,1",
			Message: "角色ID非法",
		},
	}
}
