package config

import (
	"SSPS/util"
	"fmt"

	"github.com/spf13/viper"
)

// 面板的配置文件
type PanelConfig struct {
	// 服务器ip
	ServerIp string
	// 面板端口
	PanelPort int
	// 监听时间
	ListeningTime int
	// 游戏服务器的根路径
	GameServePath string
}

// auth 配置文件读取器
var panelViper *viper.Viper

// 面板的配置实例
var PanelConf *PanelConfig

func init() {
	// 创建配置文件读取器
	panelViper = viper.New()

	// 设置配置文件
	panelViper.SetConfigName("panel")
	panelViper.SetConfigType("toml")
	panelViper.AddConfigPath("./panel_config/")
}

func init() {
	// 获取本机外网ip

	// 初始化面板实例
	PanelConf = &PanelConfig{}
	// 读取配置文件
	PanelConf.ReadPanelConfig()

	// 更新配置文件
	panelViper.Set("server_ip", util.GetExternalIP())
	panelViper.WriteConfig()
}

func (p *PanelConfig) ReadPanelConfig() *PanelConfig {

	// 读取配置文件
	err := panelViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 将配置文件映射成user
	p.ServerIp = panelViper.GetString("server_ip")
	p.PanelPort = panelViper.GetInt("panel_port")
	p.ListeningTime = panelViper.GetInt("listening_time")
	p.GameServePath = panelViper.GetString("game_serve_path")

	return p
}
