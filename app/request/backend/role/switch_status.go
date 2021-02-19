package req_role

import (
	"gf-web/library/global"
)

type SwitchStatus struct {
	Id     uint
	Status uint
}

func (r *SwitchStatus) Rules() global.Rules {
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
		{
			Field:   "status",
			Rule:    "required",
			Message: "角色状态必需",
		},
		{
			Field:   "status",
			Rule:    "in:0,1",
			Message: "角色状态非法",
		},
	}
}
