package controller

import (
	"fmt"
	"net/http"

	"SSPS/config"
	"SSPS/util"

	"github.com/gin-gonic/gin"
)

// controller 类
type ControllerLogin struct{}

var Login ControllerLogin

func init() {
	Login = ControllerLogin{}
}

// 登录回调
func (c ControllerLogin) LoginHandle(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// 创建读取 auth 配置的实例
	authStruct := config.AuthUser{}

	// 读取配置文件  获取登录账号与密码
	configUser := authStruct.ReadAuthConfig()

	// 判断账号密码是否正确
	if username != configUser.Username || password != configUser.Password {
		util.GetError().UnauthorizedError("登录失败，账号或密码错误！！！")
	}

	// 签发token
	token, err := util.UtilJWT.CreateJWT(username)
	if err != nil {
		util.GetError().ServerError(fmt.Sprintln("生成token失败，错误信息为：err: ", err))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "登录成功",
		"token": token,
	})
}
