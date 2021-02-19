package req_role

import "gf-web/library/global"

type FetchRole struct {
	Page   int    `c:"page"`
	Limit  int    `c:"limit"`
	Sort   string `c:"sort"`
	Name   string `c:"name"`
	Status []uint `c:"status"`
}

func (r *FetchRole) Rules() global.Rules {
	return global.Rules{
		{
			Field:   "page",
			Rule:    "integer",
			Message: "页码参数非法",
		},
		{
			Field:   "page",
			Rule:    "min:1",
			Message: "页码参数非法",
		},
		{
			Field:   "limit",
			Rule:    "integer",
			Message: "每页数量非法",
		},
		{
			Field:   "limit",
			Rule:    "in:10,20,30,50,100",
			Message: "每页数量非法",
		},
		{
			Field:   "sort",
			Rule:    "max-length:20",
			Message: "排序非法",
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
			Field:   "status",
			Rule:    "array:integer,0,-1,0",
			Message: "角色状态非法",
		},
	}
}
