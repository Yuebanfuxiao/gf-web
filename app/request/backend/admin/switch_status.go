package req_admin

import (
	"gf-web/library/global"
)

type SwitchStatus struct {
	Id     []uint `c:"id"`
	Status uint   `c:"status"`
}

func (r *SwitchStatus) Rules() global.Rules {
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
