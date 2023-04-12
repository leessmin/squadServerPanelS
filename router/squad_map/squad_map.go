package squad_map

import (
	"SSPS/controller"
	"SSPS/router"
)

// 地图循环

func init() {
	mapApi := router.BasicAuth.Group("/squadMap")

	// 获取地图信息
	mapApi.GET("/get", controller.SquadMap.GetSquadMap)
}
