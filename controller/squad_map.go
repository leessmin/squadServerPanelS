package controller

import (
	"SSPS/util"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// 地图循环 配置文件操作

// 地图 结构体
type controllerSquadMap struct{}

var SquadMap controllerSquadMap

func init() {
	SquadMap = controllerSquadMap{}
}

// 获取地图配置 以及地图可选列表
func (c *controllerSquadMap) GetSquadMap(ctx *gin.Context) {
	// 地图类型 Layer Level Rotation
	mapType, b := ctx.GetQuery("mapType")
	if !b {
		util.GetError().ParameterError("参数不完整")
	}

	if !(mapType == "Layer" || mapType == "Level") {
		util.GetError().ParameterError("mapType的参数只能为Layer,Level")
	}

	// 在原有的地图类型 添加Rotation.cfg  拼接成完整的文件名
	mapType = fmt.Sprintf("%vRotation.cfg", mapType)

	// 获取到已经选中的地图
	mapSelected := readSquadMap(mapType)
	// 获取可选择的地图
	mapSelect := readSquadMapBack(mapType)

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"mapSelected": mapSelected,
		"mapSelect":   mapSelect,
	}))
}

// 修改地图
func (c *controllerSquadMap) EditSquadMap(ctx *gin.Context) {
	// 定义 绑定 用户传递的json
	type bindMap struct {
		// 地图的类型
		MapType string `json:"mapType"`
		// 地图的列表
		MapList []string `json:"mapList"`
	}

	bMap := bindMap{}
	err := ctx.BindJSON(&bMap)
	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	if !(bMap.MapType == "Layer" || bMap.MapType == "Level") {
		util.GetError().ParameterError("mapType的参数只能为Layer,Level")
	}

	// 在原有的地图类型 添加Rotation.cfg  拼接成完整的文件名
	bMap.MapType = fmt.Sprintf("%vRotation.cfg", bMap.MapType)

	// 修改地图类型
	util.CreateReadWrite().InsertReplaceLineConfig(bMap.MapType, 0, strings.Join(bMap.MapList, "\n"), util.ReplaceAll{})

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"mapList": bMap.MapList,
	}))
}

// 读取地图配置
func readSquadMap(mapType string) []string {
	ch := make(chan string)

	util.CreateReadWrite().ReadNotCommentConfig(mapType, ch)

	// 储存获取到的地图列表
	var mapArr []string

	for {
		// 获取数据
		line, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		mapArr = append(mapArr, line)
	}

	return mapArr
}

// 读取备份的地图配置
func readSquadMapBack(mapType string) []string {
	// 路径拼接
	filePath := path.Join("./backCfg", mapType)

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		// 打开文件失败
		panic(fmt.Sprintf("打开文件失败,err:%v", err))
	}
	defer file.Close()

	buf := bufio.NewScanner(file)

	// 储存 地图
	var mapArr []string

	for {
		// 扫描到末尾，结束扫描
		if !buf.Scan() {
			break
		}

		// 获取当前行文字
		line := buf.Text()

		// 判断是否为注释掉的行
		if util.IsAnnotation(line) {
			// 跳过该行
			continue
		}

		// 追加
		mapArr = append(mapArr, line)
	}

	return mapArr
}
