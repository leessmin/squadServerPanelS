package login

import (
	"SSPS/controller"
	"SSPS/router"
)

// 权限 登录

func init() {
	authRouter := router.APIGroup.Group("/auth")
	{

		// 获取验证码  登录前需要调用本接口 获取验证码
		authRouter.GET("/captcha", controller.Login.CaptchaHandle)

		// 登录
		authRouter.POST("/login", controller.Login.LoginHandle)

		// 验证token是否有效
		authRouter.Any("/verify", controller.Login.VerifyTokenHandle)

		// 修改账号密码
		authRouter.POST("/updateAuth", controller.Login.UpdateAuth)
	}
}
