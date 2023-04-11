package controller

import (
	"SSPS/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// 服务器配置文件操作

// 服务器配置 结构体
type controllerSquadServer struct{}

var SquadServer controllerSquadServer

func init() {
	SquadServer = controllerSquadServer{}
}

// squad服务器配置 结构体
type serverCfg struct {
	ServerName                            string
	ShouldAdvertise                       bool
	IsLANMatch                            bool
	MaxPlayers                            int
	NumReservedSlots                      int
	PublicQueueLimit                      int
	Tags                                  string
	MapRotationMode                       string
	RandomizeAtStart                      bool
	UseVoteFactions                       bool
	UseVoteLevel                          bool
	UseVoteLayer                          bool
	AllowTeamChanges                      bool
	PreventTeamChangeIfUnbalanced         bool
	NumPlayersDiffForTeamChanges          int
	RejoinSquadDelayAfterKick             int
	RecordDemos                           bool
	AllowPublicClientsToRecord            bool
	ServerMessageInterval                 int
	ForceNonSeamlessTravelIntervalSeconds int
	TKAutoKickEnabled                     bool
	AutoTKBanNumberTKs                    int
	AutoTKBanTime                         int
	VehicleKitRequirementDisabled         bool
	AllowCommunityAdminAccess             bool
	AllowDevProfiling                     bool
	AllowQA                               bool
	VehicleClaimingDisabled               bool
}

// 获取服务器配置
func (c *controllerSquadServer) GetSquadServer(ctx *gin.Context) {
	cfg := readSquadServer()

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "获取成功", gin.H{
		"serverConfig": cfg,
	}))
}

// 修改服务器配置
func (c *controllerSquadServer) EditSquadServer(ctx *gin.Context) {
	// 接收 raw json数据
	cw, err := ctx.GetRawData()
	if err != nil {
		util.GetError().ParameterError("参数错误，请认检查参数后发送")
	}

	// 储存json数据
	mpaCfg := make(map[string]interface{})
	// 将json数据转换成map
	json.Unmarshal(cw, &mpaCfg)

	modifySquadServer(mpaCfg)

	ctx.JSON(http.StatusOK, util.CreateResponseMsg(http.StatusOK, "操作成功", gin.H{
		"config": mpaCfg,
	}))
}

// 读取服务器配置文件
func readSquadServer() serverCfg {
	ch := make(chan string)

	util.CreateReadWrite().ReadNotCommentConfig("Server.cfg", ch)

	var mapData map[string]interface{} = make(map[string]interface{})

	for {
		// 获取数据
		line, ok := <-ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		// 分割配置  取到key value
		sArr := strings.Split(line, "=")

		// 尝试转 int
		i, err := util.CreateTypeSwitch().StringToInt(sArr[1])
		if err == nil {
			// 成功
			mapData[sArr[0]] = i
			continue
		}

		// 尝试 转bool
		b, err := util.CreateTypeSwitch().StringToBool(sArr[1])
		if err == nil {
			// 成功
			mapData[sArr[0]] = b
			continue
		}

		// 都不成功
		mapData[sArr[0]] = strings.TrimSpace(sArr[1])
	}

	// 将map 转为 serverCfg
	var cfg serverCfg
	if err := mapstructure.Decode(mapData, &cfg); err != nil {
		panic(err)
	}
	return cfg
}

// 修改 服务器配置
func modifySquadServer(m map[string]interface{}) {
	// 储存需要修改的索引 和 修改的值
	var indexM map[int]string = make(map[int]string)
	// 遍历 需要修改的 map
	for key, v := range m {
		i := util.CreateReadWrite().FindContentIndex(fmt.Sprintf(`^%v.*`, key), "Server.cfg")
		indexM[i] = fmt.Sprintf(`%v=%v`, key, v)
	}

	// 遍历 indexM 修改值
	for i, v := range indexM {
		// 修改
		util.CreateReadWrite().InsertReplaceLineConfig("Server.cfg", i, v, &util.ReplaceLine{})
	}
}
