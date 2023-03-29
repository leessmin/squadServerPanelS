package middleware

import "github.com/gin-gonic/gin"

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
func NewErrorInfo(code int, msg string) *errorInfo {
	return &errorInfo{
		Code: code,
		Msg:  msg,
	}
}
