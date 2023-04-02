package system

import (
	"SSPS/controller"
	"SSPS/router"
)

// 获取系统信息

func init() {
	systemRouter := router.APIGroup.Group("/system")
	{
		// 获取系统信息 使用Websocket
		systemRouter.GET("/info", controller.System.GetSystemInfo)
	}
}
