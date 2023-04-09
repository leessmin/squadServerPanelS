package squad_admin_user

import (
	"SSPS/controller"
	"SSPS/router"
)

// 操作 管理员
func init() {
	squadAdminUser := router.BasicAuth.Group("/adminUser")

	// 获取管理员
	squadAdminUser.GET("/get", controller.SquadAdminUser.GetAdminUser)
}
