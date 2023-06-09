package controller

import (
	"SSPS/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// squad 管理组 操作

// 管理组 结构体
type controllerSquadAdminGroup struct {
}

// 管理组 实例
var SquadAdminGroup controllerSquadAdminGroup

func init() {
	SquadAdminGroup = controllerSquadAdminGroup{}
}

// 管理组 映射
type adminGroup struct {
	// 组名
	GroupName string `json:"groupName"`
	// 备注
	Info string `json:"info"`
	// 权限
	Auth []string `json:"auth"`
}

// 获取管理员组
func (c *controllerSquadAdminGroup) GetAdminGroup(ctx *gin.Context) {

	// 读取管理组
	adminGroupArr := readAdminGroup()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"adminGroup": adminGroupArr,
	}))
}

// 添加 或 编辑 管理员组
func (c *controllerSquadAdminGroup) AddEditAdminGroup(ctx *gin.Context) {

	var ag adminGroup

	err := ctx.BindJSON(&ag)
	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	// 查找 组名 是否存在
	i := util.CreateReadWrite().FindContentIndex(fmt.Sprintf(`^Group=%v:([A-z]+,{0,}){0,}([^\n]*\/\/[^\n]*){0,}`, ag.GroupName), "Admins.cfg")

	// 判断否有该组名
	if i == -1 {
		// 不存在该组名
		// 添加管理组

		// 查找管理组   找到管理组后期向管理组后面添加   未找到管理组直接在最后面添加
		ind := util.CreateReadWrite().FindContentIndex(`^Group=[A-z]*:([A-z]+,{0,}){0,}([^\n]*\/\/[^\n]*){0,}`, "Admins.cfg")

		// 判断是否已经有用户组   没有用户组  追加到文件末尾处  存在用户组  在用户组下追加
		if ind == -1 {
			// 不存在用户组
			// 追加用户组
			util.CreateReadWrite().InsertReplaceLineConfig("Admins.cfg", ind, ag.formatString(), &util.AppendLine{})
		} else {
			// 插入用户组
			util.CreateReadWrite().InsertReplaceLineConfig("Admins.cfg", ind, ag.formatString(), &util.InsertLine{})
		}
	} else {
		// 修改管理组
		util.CreateReadWrite().InsertReplaceLineConfig("Admins.cfg", i, ag.formatString(), &util.ReplaceLine{})
	}

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"adminGroup": ag,
	}))
}

// 删除 管理组
func (c *controllerSquadAdminGroup) DelAdminGroup(ctx *gin.Context) {
	groupName, b := ctx.GetQuery("groupName")
	if !b {
		util.GetError().ParameterError("参数不完整")
	}

	// 查找 组名 的索引
	index := util.CreateReadWrite().FindContentIndex(fmt.Sprintf("^Group=%v:([A-z]+,{0,}){0,}([^\\n]*\\/\\/[^\\n]*){0,}", groupName), "Admins.cfg")
	if index <= -1 {
		util.GetError().ParameterError(fmt.Sprintf("未找到组名为：“%v”的管理组", groupName))
	}

	// 删除 组名
	util.CreateReadWrite().InsertReplaceLineConfig("Admins.cfg", index, "", &util.DeleteLine{})

	// 删除 组名 下的 管理员
	util.CreateReadWrite().InsertReplaceLineConfig("Admins.cfg", 0, fmt.Sprintf(`^Admin=[0-9]*:%v`, groupName), &util.DeleteRegular{})

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{}))
}

// 读取并处理 管理员组 AdminGroup
func readAdminGroup() []adminGroup {
	ch := make(chan string)

	util.CreateReadWrite().ReadNotCommentConfig("Admins.cfg", ch)

	// 储存adminGroup
	var adminGroupArr []adminGroup

	for {
		// 获取数据
		line, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		ag := &adminGroup{}
		// 将 string 转 adminGroup
		b := ag.formatStrToAdminGroup(strings.TrimSpace(line))
		if !b {
			continue
		}

		adminGroupArr = append(adminGroupArr, *ag)
	}

	return adminGroupArr
}

// 管理组结构体格式化为相应的字符串
// 如:	Group=MyGroup: pause, demos, changemap // 备注
func (ag adminGroup) formatString() string {

	var str string

	// 判断是否有备注
	if strings.TrimSpace(ag.Info) == "" {
		// 没有备注
		str = fmt.Sprintf("Group=%v:%v", ag.GroupName, strings.Join(ag.Auth, ","))
	} else {
		str = fmt.Sprintf("Group=%v:%v // %v", ag.GroupName, strings.Join(ag.Auth, ","), ag.Info)
	}

	return str
}

// 处理 管理组 的格式  字符串转AdminGroup
// 如果字符串开头不等于Group= 则抛出异常
func (ag *adminGroup) formatStrToAdminGroup(str string) bool {

	// 判断是否符合 管理员格式
	isOk := util.CreateRegexp().VerifyStr(`^Group=[A-z]*:([A-z]+,{0,}){0,}([^\n]*\/\/[^\n]*){0,}`, str)
	if !isOk {
		return false
	}

	// 获取 管理组 名
	groupNameArr, _ := util.CreateRegexp().FindString(`(?<=Group=).*?(?=:)`, str)
	groupName := groupNameArr[0]

	// 获取备注
	infoArr, b := util.CreateRegexp().FindString(`(?<=//).*`, str)

	var auth []string
	var info string
	// 判断是否找到备注
	if b {
		// 找到备注
		info = infoArr[0]
		// 查找管理管理组的权限
		arr, _ := util.CreateRegexp().FindString(`(?<=:).*?(?=//)`, str)
		auth = strings.Split(arr[0], ",")
	} else {
		// 查找管理管理组的权限
		arr, _ := util.CreateRegexp().FindString(`(?<=:).*`, str)
		auth = strings.Split(arr[0], ",")
	}

	// 将 处理好的结果 储存到实例
	ag.GroupName = groupName
	ag.Info = info
	ag.Auth = auth

	return true
}
