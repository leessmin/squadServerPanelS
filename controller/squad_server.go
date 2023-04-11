package controller

import (
	"SSPS/util"
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
