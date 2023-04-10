package proxy

import (
	"SSPS/controller"
	"SSPS/router"
)

// 代理 外来api
func init() {
	rg := router.APIGroup.Group("/proxy")

	// 代理steam的api
	steamProxy := rg.Group("/steam")
	{
		// 获取玩家信息
		steamProxy.GET("/GetPlayerSummaries", controller.Proxy.SteamGetPlayerSummaries)
	}

}
