package controller

import (
	"SSPS/config"
	"SSPS/util"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

// controller 类
type controllerSystem struct {
	melodyWS *melody.Melody
	// 是否已经运行了sendMsg函数  true正在运行
	isSend bool
}

var System controllerSystem

func init() {
	// 创建实例
	System = controllerSystem{
		// 创建 melody实例
		melodyWS: melody.New(),
		isSend:   false,
	}

	// 连接会话时触发
	System.melodyWS.HandleConnect(func(s *melody.Session) {
		// 开启广播发送消息
		System.sendMsg()
	})
}

// 获取系统信息
func (c *controllerSystem) GetSystemInfo(ctx *gin.Context) {
	token, _ := ctx.GetQuery("token")

	// 验证token
	util.UtilJWT.VerifyToken(token)

	// 鉴权通过
	// 升级协议 WebSocket
	c.melodyWS.HandleRequest(ctx.Writer, ctx.Request)
}

// 广播 消息
func (c *controllerSystem) sendMsg() {
	// 向所有会话广播消息已经开启 防止再次开启 直接返回
	if c.isSend {
		return
	}

	// 开启广播
	c.isSend = true

	// 开启 协程 广播消息
	go func(c *controllerSystem) {

		// 关闭广播
		defer func() {
			c.isSend = false

			// 防止关闭广播期间 又新的会话建立连接
			// 关闭广播后 判断 是否存在会话  如果存在会话 重新调用广播
			if c.melodyWS.Len() > 0 {
				c.sendMsg()
			}
		}()

		forMsg(c)
	}(c)
}

// 循环发送消息
func forMsg(c *controllerSystem) {
	for {

		// 判断 是否存在会话，不存在会话,退出消息循环
		if c.melodyWS.Len() <= 0 {
			return
		}

		// 获取系统信息
		system := util.CreateSystem()
		data := system.GetSystemInfo()

		b, _ := json.Marshal(gin.H{
			"time": getNowTime(),
			"systemInfo": data,
		})

		// 发送消息
		c.melodyWS.Broadcast(b)

		// 根据配置的监听时间 来 睡眠程序
		time.Sleep(time.Duration(config.PanelConf.ListeningTime) * time.Second)
	}
}

// 获取现在的时间
func getNowTime() string {
	// 获取现在的时间
	now := time.Now()

	return now.Format("15:04:05")
}
