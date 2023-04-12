package squad_bans

import (
	"SSPS/controller"
	"SSPS/router"
)

// 封禁 玩家
func init() {
	bansApi := router.BasicAuth.Group("/bans")

	// 获取被封禁的玩家
	bansApi.GET("/get", controller.Bans.GetBansPlayer)

	// 添加 或 修改 封禁玩家
	bansApi.POST("/addEdit", controller.Bans.AddEditBansPlayer)

	// 删除 封禁玩家
	bansApi.DELETE("/del", controller.Bans.DelBansPlayer)
}
