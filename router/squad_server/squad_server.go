package squad_server

import (
	"SSPS/controller"
	"SSPS/router"
)

// squad 服务器的配置
func init() {
	serverApi := router.BasicAuth.Group("/squadServer")

	// 获取配置信息
	serverApi.GET("/get", controller.SquadServer.GetSquadServer)

	// 修改服务器配置
	serverApi.POST("/edit", controller.SquadServer.EditSquadServer)
}
