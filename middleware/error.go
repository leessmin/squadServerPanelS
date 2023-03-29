package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// 异常处理 捕获 中间件

// 异常处理中间件
func ErrorMiddlewareHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			// 捕获异常
			error := recover()

			// 异常 强制转换成 异常结构体
			info, ok := (error).(*errorInfo)

			// 判断是否 可 转换为 ErrorInfo
			// 如果可以转换，说明是主动抛出的错误
			if ok {
				// 存在错误信息

				// 发送错误信息
				ctx.JSON(info.Code, info)
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}

// 异常 结构体
type errorInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 异常 结构体 构造函数
func newErrorInfo(code int, msg string) *errorInfo {
	return &errorInfo{
		Code: code,
		Msg:  msg,
	}
}

// 错误处理 方法的 接口
type ErrorHandle interface {
	// 500 错误处理
	ServerError(msg string)
	// 404 错误
	NotFound(msg string)
	// 403 未知错误
	UnknownError(msg string)
	// 400 参数错误
	ParameterError(msg string)
	// 401 登录失败
	UnauthorizedError(msg string)
	// 403 权限认证不通过
	ForbiddenError(msg string)
	// 自定义错误  需要自行准备http code, msg
	MakeError(msg string)
}

var (
	once          sync.Once
	handling ErrorHandle
)

// 错误处理 结构体
type errorHandling struct {
}

// 单例模式 只存在一个错误处理结构体
// 获取 错误处理 结构体
func GetError() ErrorHandle {
	once.Do(func() {
		handling = &errorHandling{}
	})
	return handling
}

// 500 错误处理
func (e errorHandling) ServerError(msg string) {
	panic(newErrorInfo(http.StatusInternalServerError, msg))
}

// 404 错误
func (e errorHandling) NotFound(msg string) {
	panic(newErrorInfo(http.StatusNotFound, msg))
}

// 403 未知错误
func (e errorHandling) UnknownError(msg string) {
	panic(newErrorInfo(http.StatusForbidden, msg))
}

// 400 参数错误
func (e errorHandling) ParameterError(msg string) {
	panic(newErrorInfo(http.StatusBadRequest, msg))
}

// 401 登录失败
func (e errorHandling) UnauthorizedError(msg string) {
	panic(newErrorInfo(http.StatusUnauthorized, msg))
}

// 403 权限认证不通过
func (e errorHandling) ForbiddenError(msg string) {
	panic(newErrorInfo(http.StatusForbidden, msg))
}

// 自定义错误  需要自行准备http code, msg
func (e errorHandling) MakeError(msg string) {
	panic(newErrorInfo(http.StatusForbidden, msg))
}
