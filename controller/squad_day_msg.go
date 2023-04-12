package controller

import (
	"SSPS/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 每日消息 配置文件操作

type controllerSquadDayMsg struct{}

var SquadDayMsg controllerSquadDayMsg

func init() {
	SquadDayMsg = controllerSquadDayMsg{}
}

// 获取每日消息
func (c *controllerSquadDayMsg) GetSquadDayMsg(ctx *gin.Context) {

	dayMsg := readDayMsg()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"dayMsg": dayMsg,
	}))
}

// 编辑每日消息
func (c *controllerSquadDayMsg) EditSquadDayMsg(ctx *gin.Context) {
	// 接收 raw text数据
	dayMsg, err := ctx.GetRawData()
	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	str := string(dayMsg)

	// 写入数据
	util.CreateReadWrite().InsertReplaceLineConfig("MOTD.cfg", 0, string(str), util.ReplaceAll{})

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"dayMsg": str,
	}))
}

// 读取每日消息
func readDayMsg() (str string) {

	ch := make(chan string)
	util.CreateReadWrite().ReadConfig("MOTD.cfg", ch)

	var lineArr []string

	for {
		// 读取数据
		line, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		lineArr = append(lineArr, line)
	}

	str = strings.Join(lineArr, "\n")
	return
}
