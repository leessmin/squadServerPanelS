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

// 管理组 的实例
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
		util.GetError().ParameterError("产生错误，请认检查参数后发送")
	}

	// 查找 是否存在该组名 存在则修改
	i := util.CreateReadHandle().FindContentIndex(fmt.Sprintf("^Group=%v:([A-z]+,{0,}){0,}([^\\n]*\\/\\/[^\\n]*){0,}", ag.GroupName), "Admins.cfg")

	fmt.Println(i)
	if i == -1 {
		// TODO:修改管理组
	}else{
		// TODO:添加管理组
	}

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"obj": ag,
	}))
}

// 读取并处理 管理员组 AdminGroup
func readAdminGroup() []adminGroup {
	ch := make(chan string)

	util.CreateReadHandle().ReadConfig("Admins.cfg", ch)

	// 储存adminGroup
	var adminGroupArr []adminGroup

	for {
		// 获取数据
		data, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		// 判断前缀是否等于Group=
		// 不等于管理组说明不是与管理组有关的设置 跳过
		if data[0:6] != "Group=" {
			continue
		}

		// Group=Admin:kick,ban,changemap  // 管理员
		// 分割为 [Group, Admin:kick,ban,changemap  // 管理员]
		strArr := strings.Split(data, "=")

		// Admin:kick,ban,changemap  // 管理员
		// 分割为 [Admin:kick,ban,changemap, 管理员]
		strArr = strings.Split(strArr[1], "//")

		// 备注
		var info string
		// 判断是否存在 // 备注信息
		// 不存在备注信息，不添加备注 默认备注为""
		if len(strArr) > 1 {
			// 获取 到 备注
			info = strings.TrimSpace(strArr[1])
		}

		// 继续分割
		// Admin:kick,ban,changemap
		// 分割为 [Admin, kick,ban,changemap]
		strArr = strings.Split(strArr[0], ":")

		// 获取 管理组 组名
		groupName := strings.TrimSpace(strArr[0])

		// 继续分割
		// kick,ban,changemap
		// 分割为 [kick, ban, changemap]
		auth := strings.Split(strArr[1], ",")
		// 遍历 auth 去除两端的空格
		for key, v := range auth {
			auth[key] = strings.TrimSpace(v)
		}

		// 将 处理好的结果 储存到实例
		ag := adminGroup{
			GroupName: groupName,
			Info:      info,
			Auth:      auth,
		}

		adminGroupArr = append(adminGroupArr, ag)
	}

	return adminGroupArr
}
