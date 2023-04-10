package main

import (
	"SSPS/util"
	"net/url"
	"testing"
)

func TestProxyGet(t *testing.T) {
	hp := util.CreateHttpProxy()

	// 创建query参数
	urlQuery := url.Values{}
	// 拼接query
	urlQuery.Add("key", "1155051B4536156145C47935A24134EB")
	urlQuery.Add("steamId", "76561198816965856")
	// hp.ProxyGet(`http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=1155051B4536156145C47935A24134EB&steamids=76561198816965856`)
	hp.ProxyGet(`http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/`, urlQuery)
}
