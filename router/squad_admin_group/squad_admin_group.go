package squad_admin_group

import (
	"SSPS/controller"
	"SSPS/router"
)

// 操作 管理组

func init() {
	adminAdminGroup := router.BasicAuth.Group("/adminGroup")

	// 获取管理员组
	adminAdminGroup.GET("/getAdminGroup", controller.SquadAdminGroup.GetAdminGroup)
}
