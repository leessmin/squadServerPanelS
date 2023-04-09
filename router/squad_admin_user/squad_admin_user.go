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

	// 添加 或 修改 管理员
	squadAdminUser.POST("/addEdit", controller.SquadAdminUser.AddEditAdminUser)

	// 删除 管理组
	squadAdminUser.DELETE("/del", controller.SquadAdminUser.DelAdminUser)
}
