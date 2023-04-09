package squad_admin_group

import (
	"SSPS/controller"
	"SSPS/router"
)

// 操作 管理组

func init() {
	squadAdminGroup := router.BasicAuth.Group("/adminGroup")

	// 获取管理员组
	squadAdminGroup.GET("/get", controller.SquadAdminGroup.GetAdminGroup)

	// 添加 或 修改 管理员组
	squadAdminGroup.POST("/addEdit", controller.SquadAdminGroup.AddEditAdminGroup)

	// 删除 管理组
	squadAdminGroup.DELETE("/del", controller.SquadAdminGroup.DelAdminGroup)
}
