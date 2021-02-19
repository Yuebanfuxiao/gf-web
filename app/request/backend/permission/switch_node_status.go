package req_permission

import (
	"gf-web/library/global"
)

type SwitchNodeStatus struct {
	Id     []uint
	Status uint
}

func (r *SwitchNodeStatus) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "id",
			Rule:    "required",
			Message: "权限节点ID必需",
		},
		{
			Field:   "id",
			Rule:    "array:integer,1,-1,1",
			Message: "权限节点ID非法",
		},
		{
			Field:   "status",
			Rule:    "required",
			Message: "权限节点状态必需",
		},
		{
			Field:   "status",
			Rule:    "in:0,1",
			Message: "权限节点状态非法",
		},
	}
}
