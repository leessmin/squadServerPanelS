package controller

import (
	"SSPS/util"
	"fmt"
	"net/http"
	"strings"

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
	// 读取管理员名单
	adminUserArr := readAdminUser()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"adminUser": adminUserArr,
	}))
}

// 添加 或 编辑 管理员
func (c *controllerSquadAdminUser) AddEditAdminUser(ctx *gin.Context) {
	var au adminUser

	err := ctx.BindJSON(&au)
	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	// 查找 组名 是否存在
	i := util.CreateReadWrite().FindContentIndex(fmt.Sprintf("^Group=%v:([A-z]+,{0,}){0,}([^\\n]*\\/\\/[^\\n]*){0,}", au.GroupName), "Admins.cfg")

	if i == -1 {
		util.GetError().ParameterError(fmt.Sprintf("没有找到“%v”的管理组", au.GroupName))
	}

	// 查找是否有该管理员
	ii := util.CreateReadWrite().FindContentIndex(fmt.Sprintf(`^Admin=%v:[a-zA-Z0-9]*.*`, au.SteamId), "Admins.cfg")
	if ii == -1 {
		// 不存在
		// 添加
		util.CreateReadWrite().InsertReplaceLineConfig("Admins.cfg", 0, au.formatString(), &util.AppendLine{})
	} else {
		// 存在
		// 修改
		util.CreateReadWrite().InsertReplaceLineConfig("Admins.cfg", ii, au.formatString(), &util.ReplaceLine{})
	}

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"adminUser": au,
	}))
}

// 删除 管理员
func (c *controllerSquadAdminUser) DelAdminUser(ctx *gin.Context) {
	steamIds, b := ctx.GetQueryArray("steamIds")
	if !b {
		util.GetError().ParameterError("参数不完整")
	}

	// 储存 删除的 行 的索引
	var indexArr []int
	for _, steamId := range steamIds {
		// 查找是否有该管理员
		i := util.CreateReadWrite().FindContentIndex(fmt.Sprintf(`^Admin=%v:[a-zA-Z0-9]*.*`, steamId), "Admins.cfg")
		if i <= -1 {
			util.GetError().ParameterError(fmt.Sprintf("未找到steamId为：“%v”的管理员", steamId))
		}

		indexArr = append(indexArr, i)
	}

	// 批量删除
	for _, i := range indexArr {
		// 删除 管理员
		util.CreateReadWrite().InsertReplaceLineConfig("Admins.cfg", i, "", &util.DeleteLine{})
	}

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{}))
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
	isOk := util.CreateRegexp().VerifyStr(`^Admin=[0-9]*:[a-zA-Z0-9]*.*(\/\/[^\n]*)?`, str)
	if !isOk {
		return false
	}

	// 获取 管理员 steam id
	steamIdArr, _ := util.CreateRegexp().FindString(`(?<=Admin=).*?(?=:)`, str)
	steamId := steamIdArr[0]

	// 获取备注
	infoArr, b := util.CreateRegexp().FindString(`(?<=//).*`, str)

	var groupNameArr []string
	var info string
	// 判断是否找到备注
	if b {
		// 找到备注
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

// 管理组结构体格式化为相应的字符串
// 如:	Admin=123456:Admin // 备注
func (au adminUser) formatString() string {

	var str string

	// 判断是否有备注
	if strings.TrimSpace(au.Info) == "" {
		// 没有备注
		str = fmt.Sprintf("Group=%v:%v", au.SteamId, au.GroupName)
	} else {
		str = fmt.Sprintf("Admin=%v:%v // %v", au.SteamId, au.GroupName, au.Info)
	}

	return str
}
