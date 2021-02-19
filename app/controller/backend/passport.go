package backend

import (
	"gf-web/app/helper"
	"gf-web/app/request/backend/passport"
	"gf-web/app/service"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type Passport struct {
}

// 登录
func (c *Passport) SignIn(request *ghttp.Request) {
	var (
		req  req_passport.SignIn
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if rst, err := service.NewAdmin(request).SignIn(args); err != nil {
		g.Dump("---------")
		g.Dump(err.Error())
		g.Dump(rst)
		helper.Fail(request, err)
	} else {
		helper.Success(request, "登录成功", g.Map{
			"type":   rst.Type,
			"token":  rst.Token,
			"expire": rst.Expire,
		})
	}
}

// 注册
func (c *Passport) SignUp(request *ghttp.Request) {
	var (
		req  req_passport.SignUp
		args = make(g.Map)
	)

	helper.Accept(request, &req, args)

	if _, err := service.NewAdmin(request).SignUp(args); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "注册成功")
	}
}

// 退出
func (c *Passport) SignOut(request *ghttp.Request) {
	if err := service.NewAdmin(request).SignOut(); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "退出成功")
	}
}

// 刷新授权
func (c *Passport) RefreshAuth(request *ghttp.Request) {
	if rst, err := service.NewAdmin(request).RefreshAuth(); err != nil {
		helper.Fail(request, err)
	} else {
		helper.Success(request, "刷新成功", g.Map{
			"type":   rst.Type,
			"token":  rst.Token,
			"expire": rst.Expire,
		})
	}
}
