package login

import (
	"SSPS/controller"
	"SSPS/router"
)

// 权限 登录

func init() {
	authRouter := router.APIGroup.Group("/auth")
	{
		// 登录
		authRouter.POST("/login", controller.Login.LoginHandle)
	}
}
