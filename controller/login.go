package controller

import (
	"fmt"
	"net/http"

	"SSPS/config"
	"SSPS/util"

	"github.com/gin-gonic/gin"
)

// controller 类
type controllerLogin struct{}

var Login controllerLogin

func init() {
	Login = controllerLogin{}
}

// 获取验证码   验证码 w 70px h 35px
func (c *controllerLogin) CaptchaHandle(ctx *gin.Context) {

	// 生成验证码
	cc := util.CreateCaptcha(70, 35)

	/* ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"msg":     "验证码生成成功",
		"captcha": &cc,
	}) */

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "验证码生成成功", gin.H{
		"id":    cc.Id,
		"image": cc.Image,
	}))
}

// 登录回调
func (c *controllerLogin) LoginHandle(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	// 验证码 id
	captcha_id := ctx.PostForm("captcha_id")
	// 验证码
	captcha_code := ctx.PostForm("captcha_code")

	// 创建读取 auth 配置的实例
	authStruct := config.AuthUser{}

	// 读取配置文件  获取登录账号与密码
	configUser := authStruct.ReadAuthConfig()

	// 判断验证码
	if !util.VerifyCaptcha(captcha_id, captcha_code) {
		util.GetError().UnauthorizedError("验证码错误！！！")
	}

	// 判断账号密码是否正确
	if username != configUser.Username || password != configUser.Password {
		util.GetError().UnauthorizedError("登录失败，账号或密码错误！！！")
	}

	// 签发token
	token, err := util.UtilJWT.CreateJWT(username, configUser.Op_time)
	if err != nil {
		util.GetError().ServerError(fmt.Sprintln("生成token失败，错误信息为：err: ", err))
	}

	// 是否为初次登录  true为是
	isFirst := false
	// 判断是否是初次登录
	if len(config.CreatePanelConf().GameServePath) <= 0 || configUser.Password == "admin" {
		isFirst = true
	}

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "登录成功", gin.H{
		"token": token,
		"first": isFirst,
	}))
}

// 验证token是否有效
func (c *controllerLogin) VerifyTokenHandle(ctx *gin.Context) {
	// 从请求头中 获取token
	tokenString := ctx.GetHeader("Authorization")[7:]

	// 解析token
	claims := util.UtilJWT.VerifyToken(tokenString)

	// 创建读取 auth 配置的实例
	authStruct := config.AuthUser{}
	// 读取配置文件  获取登录账号与密码 与修改时间
	configUser := authStruct.ReadAuthConfig()

	// 判断token是否有效
	if configUser.Username != claims.Username || configUser.Op_time != claims.Op_time {
		util.GetError().ForbiddenError("token失效")
	}

	// 成功
	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "token有效", gin.H{
		"token": tokenString,
	}))
}

// 修改账号或密码
func (c *controllerLogin) UpdateAuth(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	newUsername := ctx.PostForm("newUsername")
	newPassword := ctx.PostForm("newPassword")

	// 创建读取 auth 配置的实例
	authStruct := config.AuthUser{}
	// 读取配置文件  获取登录账号与密码 与修改时间
	configUser := authStruct.ReadAuthConfig()

	// 判断账号密码是否正确
	if username != configUser.Username || password != configUser.Password {
		util.GetError().UnauthorizedError("账号或密码错误，无法完成修改")
	}

	// 判断新密码长度是否小于 8
	if len(newPassword) < 8 {
		util.GetError().ParameterError("密码长度不能小于8位")
	}

	// 判断是否有修改新账号
	if len(newUsername) > 0 {
		// 修改账号
		authStruct.Username = newUsername
	}

	// 修改密码
	authStruct.Password = newPassword

	// 更新账号密码
	authStruct.UpdateAuth()

	// 成功
	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "更新账号密码成功", gin.H{}))
}
