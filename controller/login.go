package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"SSPS/config"
	"SSPS/util"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

// controller 类
type controllerLogin struct{}

var Login controllerLogin

func init() {
	Login = controllerLogin{}
}

// 获取验证码   验证码 w 70px h 35px
func (c controllerLogin) CaptchaHandle(ctx *gin.Context) {
	ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	ctx.Writer.Header().Set("Pragma", "no-cache")

	ctx.Writer.Header().Set("Expires", "0")

	ctx.Writer.Header().Set("Content-Type", "image/png")

	id := captcha.NewLen(4)

	var content bytes.Buffer

	captcha.WriteImage(&content, id, 280, 140) //4位验证码,宽100,高50最清晰

	http.ServeContent(ctx.Writer, ctx.Request, id+".png", time.Time{}, bytes.NewReader(content.Bytes()))

	return
}

// 登录回调
func (c controllerLogin) LoginHandle(ctx *gin.Context) {
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
