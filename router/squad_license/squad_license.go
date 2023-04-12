package squad_license

import (
	"SSPS/controller"
	"SSPS/router"
)

// 服务器 认证服务器的 许可证

func init() {
	licenseApi := router.BasicAuth.Group("/license")

	// 获取许可证
	licenseApi.GET("/get", controller.SquadLicense.GetSquadLicense)

	// 编辑许可证
	licenseApi.POST("/edit", controller.SquadLicense.EditSquadLicense)
}
