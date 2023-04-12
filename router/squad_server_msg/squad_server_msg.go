package squad_server_msg

import (
	"SSPS/controller"
	"SSPS/router"
)

// 服务器消息 配置

func init() {
	serverMsgApi := router.BasicAuth.Group("/serverMsg")

	// 获取服务器消息
	serverMsgApi.GET("/get", controller.SquadServerMsg.GetSquadServerMsg)

	// 编辑服务器消息
	serverMsgApi.POST("/edit", controller.SquadServerMsg.EditSquadServerMsg)
}
