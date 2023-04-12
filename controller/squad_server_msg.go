package controller

import (
	"SSPS/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 服务器消息 操作

type controllerSquadServerMsg struct{}

var SquadServerMsg controllerSquadServerMsg

func init() {
	SquadServerMsg = controllerSquadServerMsg{}
}

// 获取服务器消息
func (c *controllerSquadServerMsg) GetSquadServerMsg(ctx *gin.Context) {

	serverMsg := readServerMsg()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"serverMsg": serverMsg,
	}))
}

// 编辑服务器消息
func (c *controllerSquadServerMsg) EditSquadServerMsg(ctx *gin.Context) {
	// 定义 绑定 用户传递的json
	type bindMsg struct {
		ServerMsg []string `json:"serverMsg"`
	}

	bMsg := bindMsg{}
	err := ctx.BindJSON(&bMsg)
	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	str := strings.Join(bMsg.ServerMsg, "\n")

	util.CreateReadWrite().InsertReplaceLineConfig("ServerMessages.cfg", 0, str, util.ReplaceAll{})

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"serverMsg": bMsg.ServerMsg,
	}))
}

// 读取 服务器消息
func readServerMsg() (msgArr []string) {
	ch := make(chan string)

	util.CreateReadWrite().ReadConfig("ServerMessages.cfg", ch)

	for {
		// 读取数据
		line, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		msgArr = append(msgArr, line)
	}

	return
}
