package controller

import (
	"SSPS/logger"
	"SSPS/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 日志

type controllerLog struct{}

var ConLog controllerLog

func init() {
	ConLog = controllerLog{}
}

// 获取日志
func (c *controllerLog) GetLog(ctx *gin.Context) {
	// 读取日志
	li := logger.CreateLogger().ReadLog()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"log": li,
	}))
}
