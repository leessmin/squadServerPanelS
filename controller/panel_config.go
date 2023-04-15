package controller

import (
	"SSPS/config"
	"SSPS/util"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 面板 配置

type controllerPanelConfig struct {
}

var PanelConfig controllerPanelConfig

func init() {
	PanelConfig = controllerPanelConfig{}
}

// 获取面板配置
func (c *controllerPanelConfig) GetPanelConfig(ctx *gin.Context) {
	// 直接获取配置
	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"config": config.CreatePanelConf(),
	}))
}

// 更新面板配置
func (c *controllerPanelConfig) UpdateConfig(ctx *gin.Context) {
	// 接收 raw json数据
	r, err := ctx.GetRawData()
	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	// 储存json数据
	mpCfg := make(map[string]interface{})
	// 将json数据转换成map
	json.Unmarshal(r, &mpCfg)

	for key, v := range mpCfg {
		config.CreatePanelConf().UpdatePanelConfig(key, v)
	}

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"config": config.CreatePanelConf(),
	}))
}
