package component

import (
	"fmt"
	"gf-web/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"log"
)

const (
	responseConfigNodeName   = "response"
	responseDefaultGroupName = "default"
)

type responseConfig struct {
	Format string // 响应格式
}

/**
 * 实例化Response
 */
func Response(name ...interface{}) *global.Response {
	var (
		group   string
		request *ghttp.Request
		config  responseConfig
	)

	for _, item := range name {
		if request == nil {
			if v, ok := item.(*ghttp.Request); ok {
				request = v
			} else if v, ok := item.(ghttp.Request); ok {
				request = &v
			}
		}

		if group == "" {
			if v, ok := item.(string); ok {
				group = v
			}
		}

		if group != "" && request != nil {
			break
		}
	}

	if group == "" {
		group = responseDefaultGroupName
	}

	configNodeKey := fmt.Sprintf("%s.%s", responseConfigNodeName, group)

	if g.Cfg().Contains(configNodeKey) == false && group == responseDefaultGroupName {
		configNodeKey = responseConfigNodeName
	}

	if err := g.Cfg().GetStruct(configNodeKey, &config); err != nil {
		log.Fatal("Response Config Error:" + err.Error())
	}

	return global.NewResponse(request, config.Format)
}
