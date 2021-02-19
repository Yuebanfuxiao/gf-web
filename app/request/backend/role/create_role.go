package req_role

import "gf-web/library/global"

type CreateRole struct {
	Name   string `c:"name"`
	Remark string `c:"remark"`
	Status uint   `c:"status"`
	Nodes  []uint `c:"nodes"`
}

func (r *CreateRole) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "name",
			Rule:    "required",
			Message: "角色名称必需",
		},
		{
			Field:   "name",
			Rule:    "min-length:2",
			Message: "角色名称至少:min个字符",
		},
		{
			Field:   "name",
			Rule:    "max-length:20",
			Message: "角色名称至多:max个字符",
		},
		{
			Field:   "remark",
			Rule:    "max-length:100",
			Message: "角色备注至多:max个字符",
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
		{
			Field:   "nodes",
			Rule:    "required",
			Message: "权限节点必需",
		},
		{
			Field:   "nodes",
			Rule:    "array:integer,1,-1,1",
			Message: "权限节点非法",
		},
	}
}
