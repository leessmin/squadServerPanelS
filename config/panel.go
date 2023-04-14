package config

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/fsnotify/fsnotify"
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

var (
	// auth 配置文件读取器
	panelViper *viper.Viper
	// 面板的配置实例
	panelConf *PanelConfig
	// 实现单例模式 使用的sync.once
	panelOnce sync.Once
)

func init() {
	// 创建配置文件读取器
	panelViper = viper.New()

	// 设置配置文件
	panelViper.SetConfigName("panel")
	panelViper.SetConfigType("toml")
	panelViper.AddConfigPath("./panel_config/")

	// 注册 监听 面板的配置 回调
	panelViper.OnConfigChange(func(e fsnotify.Event) {
		// 重新读取配置文件
		CreatePanelConf().ReadPanelConfig()
	})
	// 开启监听
	panelViper.WatchConfig()
}

func init() {
	// 获取本机外网ip

	// 读取配置文件
	CreatePanelConf().ReadPanelConfig()

	// 更新配置文件
	panelViper.Set("server_ip", getExternalIP())
	panelViper.WriteConfig()
}

// 单例模式  初始化 面板配置 实例
func CreatePanelConf() *PanelConfig {
	panelOnce.Do(func() {
		// 初始化面板实例
		panelConf = &PanelConfig{}
	})

	return panelConf
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

// 获取外部ip
func getExternalIP() string {
	res, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		panic(fmt.Sprintf("获取外部ip出错，err:%v", err))
	}
	// 关闭连接
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	// 返回外网ip
	return string(body)
}
