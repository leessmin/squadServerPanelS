package squad_day_msg

import (
	"SSPS/controller"
	"SSPS/router"
)

// 每日消息

func init() {
	dayMsgApi := router.BasicAuth.Group("/dayMsg")

	// 获取每日消息
	dayMsgApi.GET("/get", controller.SquadDayMsg.GetSquadDayMsg)

	// 编辑每日消息
	dayMsgApi.POST("/edit", controller.SquadDayMsg.EditSquadDayMsg)
}
