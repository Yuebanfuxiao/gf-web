package component

import (
	"fmt"
	gfcasbin "gf-web/library/casbin"
	"github.com/casbin/casbin/v2"
	"github.com/gogf/gf/frame/g"
	"log"
	"time"
)

const (
	appComponentNameCasbin = "gf.app.component.casbin"
	casbinConfigNodeName   = "casbin"
	casbinDefaultGroupName = "default"
)

type casbinConfig struct {
	Model     string        // 模型配置文件
	Debug     bool          // 是否开启调试模式
	Enable    bool          // 是否启用权限验证
	AutoLoad  bool          // 是否自动定期加载策略
	Duration  time.Duration // 自动加载时间间隔
	GroupName string        // 数据库组名
	TableName string        // 访问控制规则表名
}

/**
 * 实例化Casbin
 */
func Casbin(name ...string) *casbin.Enforcer {
	var (
		group  string
		config casbinConfig
	)

	group = casbinDefaultGroupName
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}

	configNodeKey := fmt.Sprintf("%s.%s", casbinConfigNodeName, group)

	if g.Cfg().Contains(configNodeKey) == false && group == casbinDefaultGroupName {
		configNodeKey = casbinConfigNodeName
	}

	if err := g.Cfg().GetStruct(configNodeKey, &config); err != nil {
		log.Fatal("Casbin Config Error:" + err.Error())
	}

	componentKey := fmt.Sprintf("%s.%s", appComponentNameCasbin, group)

	result := components.GetOrSetFuncLock(componentKey, func() interface{} {
		adapter, err := gfcasbin.NewAdapterFromOptions(&gfcasbin.Adapter{
			GroupName: config.GroupName,
			TableName: config.TableName,
		})

		if err != nil {
			log.Fatal("Casbin New Adapter Error:" + err.Error())
		}

		enforcer, err := casbin.NewEnforcer(config.Model, adapter)

		if err != nil {
			log.Fatal("Casbin Model Error:" + err.Error())
		}

		enforcer.EnableLog(config.Debug)

		enforcer.EnableEnforce(config.Enable)

		enforcer.EnableAutoNotifyWatcher(config.AutoLoad)

		return enforcer
	})

	if result != nil {
		return result.(*casbin.Enforcer)
	}

	return nil
}
