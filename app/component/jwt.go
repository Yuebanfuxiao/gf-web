package component

import (
	"fmt"
	"gf-web/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"log"
	"time"
)

const (
	appComponentNameJwt = "gf.app.component.jwt"
	jwtConfigNodeName   = "jwt"
	jwtDefaultGroupName = "default"
)

type jwtConfig struct {
	Realm          string        // 作用域
	Algorithm      string        // 算法
	Secret         string        // 密钥
	Timeout        time.Duration // TOKEN超时时间
	Refresh        time.Duration // TOKEN刷新时间
	Unique         bool          // 唯一限制
	IdentityKey    string        // 身份识别KEY
	TokenLookup    string        // TOKEN查找位置
	TokenHeadName  string        // TOKEN头名
	PublicKeyFile  string        // 公钥文件
	PrivateKeyFile string        // 私钥文件
}

/**
 * 实例化JWT
 */
func Jwt(name ...string) *global.Jwt {
	var (
		group  string
		config jwtConfig
	)

	group = jwtDefaultGroupName
	if len(name) > 0 && name[0] != "" {
		group = name[0]
	}

	configNodeKey := fmt.Sprintf("%s.%s", jwtConfigNodeName, group)

	if g.Cfg().Contains(configNodeKey) == false && group == jwtDefaultGroupName {
		configNodeKey = jwtConfigNodeName
	}

	if err := g.Cfg().GetStruct(configNodeKey, &config); err != nil {
		log.Fatal("JWT Config Error:" + err.Error())
	}

	componentKey := fmt.Sprintf("%s.%s", appComponentNameJwt, group)

	result := components.GetOrSetFuncLock(componentKey, func() interface{} {
		jwt, err := global.NewJwt(&global.Jwt{
			Realm:          config.Realm,
			Algorithm:      config.Algorithm,
			Secret:         config.Secret,
			Timeout:        config.Timeout * time.Second,
			Refresh:        config.Refresh * time.Second,
			Unique:         config.Unique,
			IdentityKey:    config.IdentityKey,
			TokenLookup:    config.TokenLookup,
			TokenHeadName:  config.TokenHeadName,
			PublicKeyFile:  config.PublicKeyFile,
			PrivateKeyFile: config.PrivateKeyFile,
		})

		if err != nil {
			glog.Fatal("JWT New Error:" + err.Error())
		}

		return jwt
	})

	if result != nil {
		return result.(*global.Jwt)
	}

	return nil
}
