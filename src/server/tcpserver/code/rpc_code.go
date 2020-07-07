package code

const (
	// rpc code
	CodeRPCSuccess                = 0
	CodeRPCFailedToGetUserInfo    = 3001
	CodeRPCFailedToVerifyPassword = 3002
	CodeRPCSystemBusy             = 3003
	CodeRPCCacheTokenExpired      = 3004
	CodeRPCUserNotExist           = 3005
)

var CodeMsg = map[int]string{
	CodeRPCFailedToGetUserInfo:    "【rpc服务】：获取用户信息失败",
	CodeRPCFailedToVerifyPassword: "【rpc服务】：密码错误",
	CodeRPCSystemBusy:             "【rpc服务】：系统发生异常，请稍后再试",
	CodeRPCCacheTokenExpired:      "【rpc服务】：token已经过期",
	CodeRPCUserNotExist:           "【rpc服务】用户不存在",
}
