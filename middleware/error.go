package middleware

import (
	"SSPS/logger"
	"SSPS/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

// 异常处理 捕获 中间件

// 异常处理中间件
func ErrorMiddlewareHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			// 捕获异常
			error := recover()

			// 异常 强制转换成 异常结构体
			info, ok := (error).(*util.ErrorInfo)

			// 判断是否 可 转换为 ErrorInfo
			// 如果可以转换，说明是主动抛出的错误
			if ok {
				// 存在错误信息

				logger.CreateLogger().Log(zapcore.WarnLevel, info.Msg)

				// 发送错误信息
				ctx.JSON(info.Code, info)
				ctx.Abort()
			} else {
				if error != nil {
					logger.CreateLogger().Log(zapcore.ErrorLevel, info.Msg)

					// 不是主动抛出的异常
					ctx.JSON(http.StatusBadRequest, gin.H{
						"code": http.StatusBadRequest,
						"msg":  fmt.Sprint("未知错误，err: ", error),
					})
					ctx.Abort()
				}
			}

		}()

		ctx.Next()
	}
}
