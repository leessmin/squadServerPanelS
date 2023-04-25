package controller

import (
	"SSPS/config"
	t_init "SSPS/init"
	"SSPS/util"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type controllerFirstInit struct{}

var FirstInit controllerFirstInit

func init() {
	FirstInit = controllerFirstInit{}
}

// 设置 新账号 密码 squad 服务器路径 初始化面板
func (c *controllerFirstInit) InitPanel(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	game_serve_path := ctx.PostForm("game_serve_path")

	// 判断是否存在 游戏配置目录
	b, _ := util.CreateReadWrite().IsDir(path.Join(game_serve_path, "SquadGame"))
	// 不存在目录
	if !b {
		util.GetError().ParameterError(fmt.Sprintf("%v,不是游戏目录，请检查后重新输入", game_serve_path))
	}

	// 判断账号和密码长度
	if len(username) < 6 || len(password) < 6 {
		util.GetError().ParameterError("账号或密码的长度不够,要6位数字才可以哦")
	}

	// 更新配置
	config.CreatePanelConf().UpdatePanelConfig(map[string]interface{}{
		"game_serve_path": game_serve_path,
	})

	// 设置账号密码
	au := config.AuthUser{
		Username: username,
		Password: password,
	}

	// 更新
	au.UpdateAuth()

	go func() {
		time.Sleep(5*time.Second)
		// 备份配置
		t_init.BackupSquadCfg()
	}()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "初始化成功", gin.H{}))
}
