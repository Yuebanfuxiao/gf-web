package req_permission

import "gf-web/library/global"

type CreateNode struct {
	Name   string `c:"name"`
	Path   string `c:"path"`
	Method string `c:"method"`
	Remark string `c:"remark"`
	Status uint   `c:"status"`
}

func (r *CreateNode) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "name",
			Rule:    "required",
			Message: "节点名称必需",
		},
		{
			Field:   "name",
			Rule:    "min-length:2",
			Message: "节点名称至少:min个字符",
		},
		{
			Field:   "name",
			Rule:    "max-length:20",
			Message: "节点名称至多:max个字符",
		},
		{
			Field:   "path",
			Rule:    "required",
			Message: "节点路径必需",
		},
		{
			Field:   "path",
			Rule:    "min-length:2",
			Message: "节点路径至少:min个字符",
		},
		{
			Field:   "path",
			Rule:    "max-length:200",
			Message: "节点路径至多:max个字符",
		},
		{
			Field:   "method",
			Rule:    "required",
			Message: "节点操作类型必需",
		},
		{
			Field:   "method",
			Rule:    "in:GET,POST,PUT,DELETE,PATCH,ALL",
			Message: "节点操作类型非法",
		},
		{
			Field:   "remark",
			Rule:    "max-length:200",
			Message: "节点备注至多:max个字符",
		},
		{
			Field:   "status",
			Rule:    "required",
			Message: "节点状态必需",
		},
		{
			Field:   "status",
			Rule:    "in:0,1",
			Message: "节点状态非法",
		},
	}
}
