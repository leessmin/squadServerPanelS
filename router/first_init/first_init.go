package first_init

import (
	"SSPS/controller"
	"SSPS/router"
)

// 初次加载服务器，需要完成的操作

func init() {
	firstRouter := router.BasicAuth.Group("/first")
	{
		// 设置 新账号 密码 squad 服务器路径
		firstRouter.POST("/init", controller.FirstInit.InitPanel)
	}
}
