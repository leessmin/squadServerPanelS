package panel_config

import (
	"SSPS/controller"
	"SSPS/router"
)

// 面板配置文件

func init() {
	panelRouter := router.BasicAuth.Group("/panel")
	{
		// 获取面板配置
		panelRouter.GET("/get", controller.PanelConfig.GetPanelConfig)

		// 更新面板配置
		panelRouter.POST("/update",controller.PanelConfig.UpdateConfig)
	}
}
