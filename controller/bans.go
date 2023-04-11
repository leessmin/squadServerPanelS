package controller

import (
	"SSPS/util"
	"fmt"
	"net/http"
	"strings"

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

// 添加 或 编辑 封禁玩家
func (c *controllerBans) AddEditBansPlayer(ctx *gin.Context) {
	var bp bansPlayer

	err := ctx.BindJSON(&bp)
	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	// 通过steamId 查找是否已经有该封禁玩家的信息
	i := util.CreateReadWrite().FindContentIndex(fmt.Sprintf(`^%v(\/\/[^\n]*)?`, bp.SteamId), "Bans.cfg")

	// 判断是否已经存在该封禁玩家
	if i == -1 {
		// 不存在
		// 添加
		util.CreateReadWrite().InsertReplaceLineConfig("Bans.cfg", 0, bp.formatString(), &util.AppendLine{})
	} else {
		// 存在 追加
		util.CreateReadWrite().InsertReplaceLineConfig("Bans.cfg", i, bp.formatString(), &util.ReplaceLine{})
	}

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"bansPlayer": bp,
	}))
}

// 删除 封禁玩家
func (c *controllerBans) DelBansPlayer(ctx *gin.Context) {
	steamIds, b := ctx.GetQueryArray("steamIds")
	if !b {
		util.GetError().ParameterError("参数不完整")
	}

	// 储存 删除的 行 的索引
	var indexArr []int
	for _, steamId := range steamIds {
		// 查找是否有该 封禁玩家
		i := util.CreateReadWrite().FindContentIndex(fmt.Sprintf(`^%v:.*`, steamId), "Bans.cfg")
		if i <= -1 {
			util.GetError().ParameterError(fmt.Sprintf("未找到steamId为：“%v”的管理员", steamId))
		}

		indexArr = append(indexArr, i)
	}

	// 删除的行数
	var delLine int = 0
	// 批量删除
	for _, i := range indexArr {
		// 删除 封禁玩家
		util.CreateReadWrite().InsertReplaceLineConfig("Bans.cfg", i-delLine, "", &util.DeleteLine{})
		delLine++
	}

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{}))
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

// 将封禁玩家结构体格式化为相应的字符串
// 如76561198039509812:0 //永久封禁-使用作弊程序
func (bp bansPlayer) formatString() string {
	var str string

	// 判断是否有备注
	if strings.TrimSpace(bp.Info) == "" {
		// 没有备注
		str = fmt.Sprintf(`%v:%v`, bp.SteamId, bp.BansTime)
	} else {
		str = fmt.Sprintf(`%v:%v // %v`, bp.SteamId, bp.BansTime, bp.Info)
	}

	return str
}
