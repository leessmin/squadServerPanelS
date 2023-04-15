package middleware

import (
	"SSPS/logger"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

// 访问接口的日志

// 中间件，获取访问的ip地址
func VisitLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取访问的客户端ip
		clientIp := ctx.ClientIP()
		// 获取请求的接口
		url := ctx.Request.URL.Path
		// 获取 请求的方法
		method := ctx.Request.Method

		// 写入日志
		logger.CreateLogger().Log(zapcore.InfoLevel, fmt.Sprintf("IP:%v |%v| %v", clientIp, method, url))

		// 放行
		ctx.Next()
	}
}
