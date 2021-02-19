package service

import (
	"gf-web/app/errcode"
	"gf-web/app/helper"
	"gf-web/app/model/admin"
	"gf-web/app/model/admin_login_record"
	"gf-web/app/model/role"
	"gf-web/app/repository"
	"gf-web/library/global"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"net/http"
)

type Admin struct {
	Abstract
	admin            *repository.Admin
	adminLoginRecord *repository.AdminLoginRecord
}

// 创建管理员参数
type SignInArgs struct {
	Account  string
	Email    string
	Mobile   string
	Password string
}

// 创建管理员参数
type SignUpArgs struct {
	Account  string
	Email    string
	Mobile   string
	Password string
}

// 创建管理员参数
type CreateAdminArgs struct {
	Account  string
	Email    string
	Mobile   string
	Password string
	Nickname string
	Avatar   string
	Status   uint
}

// 修改管理员参数
type UpdateAdminArgs struct {
	Nickname string
	Avatar   string
	Status   uint
}

// 新建管理员服务
func NewAdmin(request ...*ghttp.Request) *Admin {
	service := &Admin{}

	if len(request) > 0 {
		service.WithRequest(request[0])
	}

	service.admin = repository.NewAdmin(service.request)

	service.adminLoginRecord = repository.NewAdminLoginRecord()

	return service
}

// 管理员注册
func (s *Admin) SignUp(data g.Map) (*admin.Entity, error) {
	data["status"] = admin.StatusDisabled

	return s.CreateAdmin(data)
}

// 管理员登录
func (s *Admin) SignIn(data g.Map) (*global.Token, error) {
	var (
		entity *admin.Entity
		err    error
	)

	if account, ok := data["account"]; ok && account != "" {
		entity, err = s.GetByAccount(gconv.String(account))

		if err != nil {
			return nil, err
		}
	} else if email, ok := data["email"]; ok && email != "" {
		entity, err = s.GetByEmail(gconv.String(email))

		if err != nil {
			return nil, err
		}
	} else if mobile, ok := data["mobile"]; ok && mobile != "" {
		entity, err = s.GetByMobile(gconv.String(mobile))

		if err != nil {
			return nil, err
		}
	} else {
		return nil, global.NewError(http.StatusBadRequest, errcode.CodeParamsInvalid, "无效参数")
	}

	if _, err := helper.PasswordCompare(entity.Password, gconv.String(data["password"])); err != nil {
		return nil, global.NewError(http.StatusBadRequest, errcode.CodePassportPasswordNotMatch, "账号或密码错误", err)
	} else {
		if entity.Status != admin.StatusEnabled {
			return nil, global.NewError(http.StatusForbidden, errcode.CodeAdminDisabled, "管理员已禁用", err)
		}

		entity, err = s.admin.Update(nil, entity, g.Map{
			"last_login_at": gtime.Now(),
			"last_login_ip": s.request.GetClientIp(),
		})

		if err != nil {
			return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
		}

		if _, err := s.WriteLoginRecord(entity.Id); err != nil {
			return nil, err
		}

		return s.GenerateAuth(entity)
	}
}

// 管理员登出
func (s *Admin) SignOut() error {
	return s.DestroyAuth()
}

// 生成授权
func (s *Admin) GenerateAuth(admin *admin.Entity) (*global.Token, error) {
	var (
		err      error
		entities []*role.Entity
		token    *global.Token
		roles    = make(g.List, 0)
	)

	entities, err = NewRole(s.request).GetRolesByAdminId(admin.Id)

	if err != nil {
		return nil, err
	}

	for _, entity := range entities {
		roles = append(roles, g.Map{
			"id":   entity.Id,
			"name": entity.Name,
		})
	}

	token, err = helper.Jwt(helper.JwtGroupBackend).GenerateToken(g.Map{
		"id":      admin.Id,
		"account": admin.Account,
		"email":   admin.Email,
		"mobile":  admin.Mobile,
		"roles":   roles,
	})

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeServerException, "服务异常", err)
	}

	return token, nil
}

/**
 * 刷新授权
 */
func (s *Admin) RefreshAuth() (*global.Token, error) {
	if token, err := helper.Jwt(helper.JwtGroupBackend).RefreshToken(s.request); err != nil {
		if e, ok := err.(*global.JwtError); ok {
			switch e.Errors {
			case
				global.JwtErrorInvalidToken,
				global.JwtErrorEmptyAuthHeader,
				global.JwtErrorEmptyToken,
				global.JwtErrorInvalidAuthHeader,
				global.JwtErrorEmptyQueryToken,
				global.JwtErrorEmptyCookieToken,
				global.JwtErrorEmptyParamToken:
				return nil, global.NewError(http.StatusUnauthorized, errcode.CodeUnauthorized, "未授权", err)
			case global.JwtErrorAuthorizeElsewhere:
				return nil, global.NewError(http.StatusUnauthorized, errcode.CodeAuthorizeElsewhere, "账号在其它地方登陆", err)
			}
		}

		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeServerException, "服务异常", err)
	} else {
		return token, nil
	}
}

// 销毁授权
func (s *Admin) DestroyAuth(ids ...uint) error {
	jwt := helper.Jwt(helper.JwtGroupBackend)

	if len(ids) > 0 {
		for _, id := range ids {
			if err := jwt.DelUniqueIdentificationCode(id); err != nil {
				continue
			}
		}
	} else {
		if err := jwt.DestroyToken(s.request); err != nil {
			if e, ok := err.(*global.JwtError); ok {
				switch e.Errors {
				case global.JwtErrorFailedTokenDestroy:
					return global.NewError(http.StatusInternalServerError, errcode.CodeServerException, "服务异常", err)
				}
			}
		}
	}

	return nil
}

// 创建管理员
func (s *Admin) CreateAdmin(data g.Map) (*admin.Entity, error) {
	if account, ok := data["account"]; ok {
		if exist, err := s.CheckAccountExist(account.(string)); err != nil {
			return nil, err
		} else {
			if exist {
				return nil, global.NewError(http.StatusBadRequest, errcode.CodeAccountExist, "账号已被占用")
			}
		}
	}

	if email, ok := data["email"]; ok {
		if exist, err := s.CheckEmailExist(email.(string)); err != nil {
			return nil, err
		} else {
			if exist {
				return nil, global.NewError(http.StatusBadRequest, errcode.CodeEmailExist, "邮箱已被占用")
			}
		}
	}

	if mobile, ok := data["mobile"]; ok {
		if exist, err := s.CheckMobileExist(mobile.(string)); err != nil {
			return nil, err
		} else {
			if exist {
				return nil, global.NewError(http.StatusBadRequest, errcode.CodeEmailExist, "手机号已被占用")
			}
		}
	}

	if password, err := helper.PasswordEncrypt(gconv.String(data["password"])); err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeServerException, "服务异常", err)
	} else {
		data["password"] = password

		if entity, err := s.admin.Create(nil, data); err != nil {
			return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
		} else {
			return entity, nil
		}
	}
}

// 删除管理员
func (s *Admin) DeleteAdmin(id uint) (int64, error) {
	return s.BatchDeleteAdmin(id)
}

// 更新管理员
func (s *Admin) UpdateAdmin(id uint, data g.Map) (*admin.Entity, error) {
	var (
		entity *admin.Entity
		err    error
	)

	entity, err = s.GetById(id)

	if err != nil {
		return nil, err
	}

	if _, err := s.admin.Update(nil, entity, data); err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	if status, ok := data["status"]; ok {
		if status == admin.StatusDisabled {
			_ = s.DestroyAuth(id)
		}
	}

	return nil, nil
}

// 批量更新管理员
func (s *Admin) BatchUpdateAdmin(data g.Map, ids ...uint) (int64, error) {
	rows, err := s.admin.BatchUpdate(nil, data, "id in(?)", ids)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	if status, ok := data["status"]; ok {
		if status == admin.StatusDisabled {
			_ = s.DestroyAuth(ids...)
		}
	}

	return rows, nil
}

// 批量删除管理员
func (s *Admin) BatchDeleteAdmin(ids ...uint) (int64, error) {
	var (
		err        error
		rows       int64
		permission = NewPermission(s.request)
	)

	_, err = permission.DeleteGroupingPoliciesByUserIds(ids...)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	rows, err = s.admin.BatchDelete(nil, "id in(?)", ids)

	if err != nil {
		return 0, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	_ = s.DestroyAuth(ids...)

	return rows, nil
}

// 批量切换管理员状态
func (s *Admin) BatchSwitchStatus(status uint, ids ...uint) (int64, error) {
	return s.BatchUpdateAdmin(g.Map{"status": status}, ids...)
}

// 重置密码
func (s *Admin) ResetPassword(id uint, password string) (*admin.Entity, error) {
	entity, err := s.GetById(id)

	if err != nil {
		return nil, err
	}

	if password, err := helper.PasswordEncrypt(password); err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeServerException, "服务异常", err)
	} else {
		entity, err = s.UpdateAdmin(id, g.Map{"password": password})

		if err != nil {
			return nil, err
		}

		_ = s.DestroyAuth(entity.Id)

		return entity, nil
	}
}

// 修改密码
func (s *Admin) UpdatePassword(id uint, oldPassword string, newPassword string) (*admin.Entity, error) {
	entity, err := s.GetById(id)

	if err != nil {
		return nil, err
	}

	if _, err := helper.PasswordCompare(entity.Password, oldPassword); err != nil {
		return nil, global.NewError(http.StatusBadRequest, errcode.CodePassportPasswordNotMatch, "旧密码错误", err)
	}

	if password, err := helper.PasswordEncrypt(newPassword); err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeServerException, "服务异常", err)
	} else {
		entity, err = s.UpdateAdmin(id, g.Map{"password": password})

		if err != nil {
			return nil, err
		}

		_ = s.DestroyAuth(id)

		return entity, nil
	}
}

// 根据ID获取管理员
func (s *Admin) GetById(id uint) (*admin.Entity, error) {
	entity, err := s.admin.GetById(id)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	if entity == nil {
		return nil, global.NewError(http.StatusNotFound, errcode.CodeAdminNotExist, "管理员不存在")
	}

	return entity, nil
}

// 根据账号获取管理员
func (s *Admin) GetByAccount(account string) (*admin.Entity, error) {
	entity, err := s.admin.Get("account = ?", account)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	if entity == nil {
		return nil, global.NewError(http.StatusNotFound, errcode.CodeAdminNotExist, "管理员不存在")
	}

	return entity, nil
}

// 根据邮箱获取管理员
func (s *Admin) GetByEmail(email string) (*admin.Entity, error) {
	entity, err := s.admin.Get("email = ?", email)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	if entity == nil {
		return nil, global.NewError(http.StatusNotFound, errcode.CodeAdminNotExist, "管理员不存在")
	}

	return entity, nil
}

// 根据手机号获取管理员
func (s *Admin) GetByMobile(mobile string) (*admin.Entity, error) {
	entity, err := s.admin.Get("mobile = ?", mobile)

	if err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
	}

	if entity == nil {
		return nil, global.NewError(http.StatusNotFound, errcode.CodeAdminNotExist, "管理员不存在")
	}

	return entity, nil
}

// 检测账号是否存在
func (s *Admin) CheckAccountExist(account string) (bool, error) {
	if account != "" {
		if exist, err := s.admin.Exist("account", account); err != nil {
			return false, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
		} else {
			if exist {
				return true, nil
			}
		}
	}

	return false, nil
}

// 检测邮箱是否存在
func (s *Admin) CheckEmailExist(email string) (bool, error) {
	if email != "" {
		if exist, err := s.admin.Exist("email", email); err != nil {
			return false, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
		} else {
			if exist {
				return true, nil
			}
		}
	}

	return false, nil
}

// 检测手机号是否存在
func (s *Admin) CheckMobileExist(mobile string) (bool, error) {
	if mobile != "" {
		if exist, err := s.admin.Exist("mobile", mobile); err != nil {
			return false, global.NewError(http.StatusInternalServerError, errcode.CodeDbQueryException, "服务异常", err)
		} else {
			if exist {
				return true, nil
			}
		}
	}

	return false, nil
}

// 写入登录日志
func (s *Admin) WriteLoginRecord(adminId uint) (*admin_login_record.Entity, error) {
	entity := &admin_login_record.Entity{
		AdminId: adminId,
		LoginAt: gtime.Now(),
		LoginIp: s.request.GetClientIp(),
	}

	if _, err := s.adminLoginRecord.Insert(entity); err != nil {
		return nil, global.NewError(http.StatusInternalServerError, errcode.CodeDbExecException, "服务异常", err)
	}

	return entity, nil
}
