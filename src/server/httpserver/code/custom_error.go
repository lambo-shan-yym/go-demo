package code

import (
	"fmt"
	"net/http"
)

type Exception struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
}
const (
	// http code
	CodeSuccess = 0
)


var (
	Success         = NewException(http.StatusOK, CodeSuccess, "")
	NotFound        = NewException(http.StatusNotFound, 200404, "找不到该路由")
	NoMethod        = NewException(http.StatusNotFound, 200405, "找不到该方法")
	ServerException = NewException(http.StatusInternalServerError, 200500, "系统未知异常,请稍后再试！")
	ParamInvalid    = NewException(http.StatusBadRequest, 200400, "参数不合法：%s")
	RequiredParam   = NewException(http.StatusBadRequest, 200401, "缺少参数：%s")
	Unauthorized    = NewException(http.StatusUnauthorized, 200410, "请求未携带token，无权限访问!")
	CreateTokenFail = NewException(http.StatusOK, 200411, "token生成失败")
	TokenExpired    = NewException(http.StatusUnauthorized, 200412, "授权已过期")
	SystemException = NewException(http.StatusInternalServerError, 200501, "系统异常:%s")
)

func (e *Exception) Error() string {
	return e.Msg
}

func OtherException(message string) *Exception {
	return NewException(http.StatusForbidden, 100403, message)
}

func NewException(statusCode int, code int, msg string) *Exception {
	return &Exception{
		StatusCode: statusCode,
		Code:       code,
		Msg:        msg,
	}
}

func FormatException(e *Exception, msg string) *Exception {
	exception := *e
	exception.Msg = fmt.Sprintf(exception.Msg, msg)
	return &exception
}
