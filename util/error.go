package util

import (
	"net/http"
	"sync"
)

// 抛出异常 处理

// 异常 结构体
type ErrorInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 异常 结构体 构造函数
func NewErrorInfo(code int, msg string) *ErrorInfo {
	return &ErrorInfo{
		Code: code,
		Msg:  msg,
	}
}

// 错误处理 方法的 接口
type errorHandle interface {
	// 500 错误处理
	ServerError(msg string)
	// 404 错误
	NotFound(msg string)
	// 400 参数错误
	ParameterError(msg string)
	// 401 登录失败
	UnauthorizedError(msg string)
	// 403 权限认证不通过
	ForbiddenError(msg string)
	// 自定义错误  需要自行准备http code, msg
	MakeError(code int, msg string)
}

var (
	onceErr     sync.Once
	handling errorHandle
)

// 错误处理 结构体
type errorHandling struct {
}

// 单例模式 只存在一个错误处理结构体
// 获取 错误处理 结构体
func GetError() errorHandle {
	onceErr.Do(func() {
		handling = &errorHandling{}
	})
	return handling
}

// 500 错误处理
func (e *errorHandling) ServerError(msg string) {
	panic(NewErrorInfo(http.StatusInternalServerError, msg))
}

// 404 错误
func (e *errorHandling) NotFound(msg string) {
	panic(NewErrorInfo(http.StatusNotFound, msg))
}

// 400 参数错误
func (e *errorHandling) ParameterError(msg string) {
	panic(NewErrorInfo(http.StatusBadRequest, msg))
}

// 401 登录失败
func (e *errorHandling) UnauthorizedError(msg string) {
	panic(NewErrorInfo(http.StatusUnauthorized, msg))
}

// 403 权限认证不通过
func (e *errorHandling) ForbiddenError(msg string) {
	panic(NewErrorInfo(http.StatusForbidden, msg))
}

// 自定义错误  需要自行准备http code, msg
func (e *errorHandling) MakeError(code int, msg string) {
	panic(NewErrorInfo(http.StatusForbidden, msg))
}
