package req_permission

import (
	"gf-web/library/global"
)

type FetchNode struct {
	Page  int `c:"page"`
	Limit int `c:"limit"`
}

func (r *FetchNode) Rules() global.Rules {
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
	}
}
