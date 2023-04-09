package controller

import (
	"SSPS/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// squad 管理员 操作

// 管理员 结构体
type controllerSquadAdminUser struct{}

// 管理员实例
var SquadAdminUser controllerSquadAdminUser

func init() {
	SquadAdminUser = controllerSquadAdminUser{}
}

// 管理员 映射
type adminUser struct {
	// 所属 管理组 组名
	GroupName string `json:"groupName"`
	// 备注
	Info string `json:"info"`
	// steam id
	SteamId string `json:"steamId"`
}

// 获取管理员
func (c *controllerSquadAdminUser) GetAdminUser(ctx *gin.Context) {
	// readAdminUser()

	// 读取管理员名单
	adminUserArr := readAdminUser()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"adminUser": adminUserArr,
	}))
}

// 读取管理员
func readAdminUser() []adminUser {
	ch := make(chan string)

	util.CreateReadWrite().ReadNotCommentConfig("Admins.cfg", ch)

	// 储存adminUser
	var adminUserArr []adminUser

	for {
		// 获取数据
		line, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		au := &adminUser{}
		// 将 获取到的字符串 转 adminUser
		b := au.formatStrToAdminUser(line)
		// 判断转换是否成功
		if !b {
			// 不成功 跳出此次循环
			continue
		}

		adminUserArr = append(adminUserArr, *au)
	}

	return adminUserArr
}

// ^Admin=[0-9]*:[a-zA-Z0-9]*.*(\/\/[^\n]*)

// 处理 管理员格式   字符串转adminUser
// 不符合条件的直接 返回false
func (au *adminUser) formatStrToAdminUser(str string) bool {

	// 判断是否符合 管理员格式
	isOk := util.CreateRegexp().VerifyStr(`^Admin=[0-9]*:[a-zA-Z0-9]*.*(\/\/[^\n]*)`, str)
	if !isOk {
		return false
	}

	// 获取steam id
	steamIdArr, _ := util.CreateRegexp().FindString(`(?<=Admin=).*?(?=:)`, str)
	steamId := steamIdArr[0]

	// 获取备注
	infoArr, b := util.CreateRegexp().FindString(`(?<=//).*`, str)

	var groupNameArr []string
	var info string
	// 判断是否找到注释
	if b {
		// 找到注释
		info = infoArr[0]
		// 查找管理员所属的用户组
		groupNameArr, _ = util.CreateRegexp().FindString(`(?<=:).*?(?=//)`, str)
	} else {
		// 查找管理员所属的用户组
		groupNameArr, _ = util.CreateRegexp().FindString(`(?<=:).*`, str)
	}
	groupName := groupNameArr[0]

	// 赋值
	au.SteamId = steamId
	au.Info = info
	au.GroupName = groupName

	return true
}
