package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Abstract struct {
	request *ghttp.Request
}

// 分页
type Pagination struct {
	Page  int     `json:"page"`
	Total int     `json:"total"`
	Limit int     `json:"limit"`
	List  []g.Map `json:"list"`
}

// 携带请求
func (s *Abstract) WithRequest(request *ghttp.Request) {
	s.request = request
}
