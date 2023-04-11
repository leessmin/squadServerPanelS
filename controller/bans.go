package controller

import (
	"SSPS/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// squad 封禁玩家

// 封禁玩家 结构体
type controllerBans struct{}

// 封禁玩家 实例
var Bans controllerBans

func init() {
	Bans = controllerBans{}
}

// 封禁的玩家
type bansPlayer struct {
	// 封禁玩家的steam id
	SteamId string `json:"steamId"`
	// 封禁的时间
	BansTime string `json:"bansTime"`
	// 备注
	Info string `json:"info"`
}

// 获取封禁的玩家
func (c *controllerBans) GetBansPlayer(ctx *gin.Context) {
	// 读取封禁玩家的名单
	bansPlayerArr := readBansPlayer()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"bansPlayers": bansPlayerArr,
	}))

}

// 读取封禁玩家的名单
func readBansPlayer() []bansPlayer {

	ch := make(chan string)

	util.CreateReadWrite().ReadNotCommentConfig("Bans.cfg", ch)

	// 储存 封禁的玩家
	var bansPlayerArr []bansPlayer

	for {
		// 读取数据
		line, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		bp := &bansPlayer{}
		// 将 获取到的字符串转 bansPlayer
		b := bp.formatStrToBansPlayer(line)
		// 判断转换是否成功
		if !b {
			// 不成功 跳出循环
			continue
		}

		bansPlayerArr = append(bansPlayerArr, *bp)

	}

	return bansPlayerArr

}

// 处理 封禁玩家格式   字符串转bansPlayer
// 不符合条件的直接 返回false

func (bp *bansPlayer) formatStrToBansPlayer(str string) bool {

	// 判断是否符合 封禁玩家格式
	isOk := util.CreateRegexp().VerifyStr(`^[0-9]*:[0-9]*.*(\/\/[^\n]*)?`, str)
	if !isOk {
		return false
	}

	// 获取 封禁玩家 steam id
	steamIdArr, _ := util.CreateRegexp().FindString(`[0-9].*?(?=:)`, str)
	steamId := steamIdArr[0]

	// 获取备注
	// 获取备注
	infoArr, b := util.CreateRegexp().FindString(`(?<=//).*`, str)

	var info string
	var bansTime string
	// 判断是否有备注
	if b {
		// 找到备注
		info = infoArr[0]
		// 封禁玩家的封禁时间
		bansTimeArr, _ := util.CreateRegexp().FindString(`(?<=:).*?(?=//)`, str)
		bansTime = bansTimeArr[0]
	} else {
		// 封禁玩家的封禁时间
		bansTimeArr, _ := util.CreateRegexp().FindString(`(?<=:).*`, str)
		bansTime = bansTimeArr[0]
	}

	// 赋值
	bp.SteamId = steamId
	bp.Info = info
	bp.BansTime = bansTime

	return true
}
