package req_permission

import (
	"gf-web/library/global"
)

type DeleteNode struct {
	Id []uint `c:"id"`
}

func (r *DeleteNode) Rules() global.Rules {
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
	}
}
