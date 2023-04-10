package controller

import (
	"SSPS/util"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// 代理 外来api
type controllerProxy struct{}

// 代理实例
var Proxy controllerProxy

func init() {
	Proxy = controllerProxy{}
}

// 获取玩家信息
func (c *controllerProxy) SteamGetPlayerSummaries(ctx *gin.Context) {
	// 获取query参数
	key := ctx.Query("key")
	steamIds := ctx.Query("steamIds")

	// 创建query参数
	urlQuery := url.Values{}
	// 拼接query
	urlQuery.Add("key", key)
	urlQuery.Add("steamIds", steamIds)

	hp := util.CreateHttpProxy()
	// 向steam api发送请求获取steamId下的数据
	s := hp.ProxyGet(`http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/`, urlQuery)

	// 将获取到的数据转json
	var i interface{}
	json.Unmarshal(s, &i)

	ctx.JSON(http.StatusOK, i)
}
