package bans

import (
	"SSPS/controller"
	"SSPS/router"
)

// 封禁 玩家
func init() {
	bansApi := router.BasicAuth.Group("/bans")

	// 获取被封禁的玩家
	bansApi.GET("/get", controller.Bans.GetBansPlayer)
}
