package errcode

const (
	// 所有错误码定义遵循以下规则
	// 前三位为http状态码，后三位为自定义错误码
	// 自定义错误码应从100起开始定义，预留出前100的错误码位

	// 参数无效
	CodeParamsInvalid = 400000
	// 账号或密码错误
	CodePassportPasswordNotMatch = 400101
	// 账号已经存在
	CodeAccountExist = 400102
	// 邮箱已经存在
	CodeEmailExist = 400103
	// 手机号已经存在
	CodeMobileExist = 400104
	// 创建失败
	CodeCreateFailed = 400105
	// 更新失败
	CodeUpdateFailed = 400106
	// 删除失败
	CodeDeleteFailed = 400107
	// 权限节点已经存在
	CodePermissionNodeExist = 400108

	// 未授权
	CodeUnauthorized = 401101
	// 授权已过期
	CodeTokenExpired = 401102
	// 无效授权
	CodeTokenInvalid = 401103
	// 账号在其它地方登陆
	CodeAuthorizeElsewhere = 401104

	// 无权限
	CodeNoPermission = 403101
	// 管理员已禁用
	CodeAdminDisabled = 403102

	// 管理员不存在
	CodeAdminNotExist = 404101
	// 账号不存在
	CodeAccountNotExist = 404102
	// 邮箱不存在
	CodeEmailNotExist = 404103
	// 手机号不存在
	CodeMobileNotExist = 404104

	// 参数验证失败
	CodeParameterVerificationFailed = 422101

	// 数据库查询失败
	CodeDbQueryException = 500101
	// 数据库执行失败
	CodeDbExecException = 500102
	// 服务内部异常
	CodeServerException = 500103
)
