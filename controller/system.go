package controller

import (
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

	// 升级协议 WebSocket
	c.melodyWS.HandleRequest(ctx.Writer, ctx.Request)
}

// 广播 循环 发送消息
func (c *controllerSystem) sendMsg() {
	// 向所有会话广播消息已经开启   直接返回
	if c.isSend {
		return
	}

	// 开启广播
	c.isSend = true

	// 开启 协程处理
	go func(c *controllerSystem) {

		// 关闭广播
		defer func() {
			c.isSend = false
		}()

		for {
			// 获取所有的会话
			sessions, _ := c.melodyWS.Sessions()

			// 判断 是否存在会话，不存在会话,退出消息循环
			if len(sessions) <= 0 {
				return
			}

			// TODO: 发送消息
			// 获取系统信息再发送消息
			c.melodyWS.Broadcast([]byte("你好呀"))

			// TODO: 暂时 写死 5秒
			time.Sleep(time.Second * 5)
		}
	}(c)

}
