package controller

import (
	"SSPS/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type controllerSquadLicense struct{}

var SquadLicense controllerSquadLicense

func init() {
	SquadLicense = controllerSquadLicense{}
}

// 获取服务器 许可证
func (c *controllerSquadLicense) GetSquadLicense(ctx *gin.Context) {
	license := readLicense()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"license": license,
	}))
}

// 编辑 服务器 许可证
func (c *controllerSquadLicense) EditSquadLicense(ctx *gin.Context) {
	// 接收 raw text数据
	license, err := ctx.GetRawData()

	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	str := string(license)

	// 写入数据
	util.CreateReadWrite().InsertReplaceLineConfig("License.cfg", 0, string(str), util.ReplaceAll{})

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"license": str,
	}))
}

// 读取服务器许可证
func readLicense() (str string) {
	ch := make(chan string)
	util.CreateReadWrite().ReadConfig("License.cfg", ch)

	var lineArr []string

	for {
		// 读取数据
		line, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		lineArr = append(lineArr, line)
	}

	str = strings.Join(lineArr, "\n")
	return
}
