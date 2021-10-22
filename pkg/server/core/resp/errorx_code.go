package resp

// 错误码
const (
	Success = 0

	ServerError = 10000
	ParamError  = 10001

	PasswordDecryptError    = 10010
	UsernameOrPasswordError = 10011
	GetUserInfoError        = 10012
	ParamTokenError         = 10013
	TokenVerifyError        = 10014
	TokenExpireError        = 10015
)

var errMap = map[int]string{
	Success: "成功",

	ServerError: "服务器内部错误",
	ParamError:  "参数错误",

	PasswordDecryptError:    "密码解析错误",
	UsernameOrPasswordError: "用户名或密码错误",
	GetUserInfoError:        "获取用户信息失败",
	ParamTokenError:         "参数token参数缺失",
	TokenVerifyError:        "token校验失败",
	TokenExpireError:        "token过期",
}
