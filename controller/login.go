package controller

import (
	"net/http"

	"SSPS/config"

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
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code": http.StatusUnauthorized,
			"msg":  "账号或密码错误，请重试",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
	})
}
