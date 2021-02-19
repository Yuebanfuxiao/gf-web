package router

import (
	"gf-web/app/controller/backend"
	"gf-web/app/middleware"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	passport   = &backend.Passport{}
	admin      = &backend.Admin{}
	role       = &backend.Role{}
	permission = &backend.Permission{}
)

func init() {
	s := g.Server()

	s.Group("/backend", func(group *ghttp.RouterGroup) {
		group.Middleware(middleware.Cors)

		group.Group("/passport", func(group *ghttp.RouterGroup) {
			// 登录账号
			group.POST("/sign-in", passport.SignIn)
			// 注册账号
			group.POST("/sign-up", passport.SignUp)
			// 退出账号
			group.DELETE("/sign-out", passport.SignOut)
			// 刷新授权
			group.POST("/refresh-auth", passport.RefreshAuth)
		})

		group.Group("/", func(group *ghttp.RouterGroup) {
			//group.Middleware(middleware.JwtAuthBackend)
			//group.Middleware(middleware.RbacBackend)

			group.Group("/admin", func(group *ghttp.RouterGroup) {
				// 创建管理员
				group.POST("/create-admin", admin.CreateAdmin)
				// 删除管理员
				group.DELETE("/delete-admin", admin.DeleteAdmin)
				// 修改管理员
				group.PUT("/update-admin", admin.UpdateAdmin)
				// 切换状态
				group.PUT("/switch-status", admin.SwitchStatus)
				// 修改密码
				group.PUT("/update-password", admin.UpdatePassword)
				// 重置密码
				group.PUT("/reset-password", admin.ResetPassword)
			})

			// 角色管理
			group.Group("/role", func(group *ghttp.RouterGroup) {
				// 拉取角色
				group.GET("/fetch-role", role.FetchRole)
				// 创建角色
				group.POST("/create-role", role.CreateRole)
				// 修改角色
				group.PUT("/update-role", role.UpdateRole)
				// 删除角色
				group.DELETE("/delete-role", role.DeleteRole)
				// 切换状态
				group.PUT("/switch-status", role.SwitchStatus)
			})

			// 节点管理
			group.Group("/permission", func(group *ghttp.RouterGroup) {
				// 节点列表
				group.GET("/fetch-node", permission.FetchNode)
				// 创建节点
				group.POST("/create-node", permission.CreateNode)
				// 修改节点
				group.PUT("/update-node", permission.UpdateNode)
				// 删除节点
				group.DELETE("/delete-node", permission.DeleteNode)
				// 切换状态
				group.PUT("/switch-node-status", permission.SwitchNodeStatus)
			})
		})
	})
}
