package logger

import (
	"SSPS/controller"
	"SSPS/router"
)

// 日志 处理

func init() {
	logRouter := router.BasicAuth.Group("/log")
	{
		// 获取日志
		logRouter.GET("/get", controller.ConLog.GetLog)
	}
}
