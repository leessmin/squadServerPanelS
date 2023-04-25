package config

import (
	"fmt"
	"math"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 面板的配置文件
type PanelConfig struct {
	// 服务器ip
	ServerIp string `json:"server_ip"`
	// 面板端口
	PanelPort int `json:"panel_port"`
	// 监听时间
	ListeningTime int `json:"listening_time"`
	// 游戏服务器的根路径
	GameServePath string `json:"game_serve_path"`
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

	// 获取ip
	ip, err := getExternalIP()
	if err != nil {
		panic(err)
	}

	// 更新配置文件
	panelViper.Set("server_ip", ip)
	panelViper.WriteConfig()
}

func init() {
	// 设置前端请求页面的地址
	setWebPort(panelConf.ServerIp, panelConf.PanelPort)
}

// 单例模式  初始化 面板配置 实例
func CreatePanelConf() *PanelConfig {
	panelOnce.Do(func() {
		// 初始化面板实例
		panelConf = &PanelConfig{}
	})

	return panelConf
}

// 读取配置
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

// 更新配置
func (p *PanelConfig) UpdatePanelConfig(mapValue map[string]interface{}) error {

	for key, value := range mapValue {
		// 判断是否 存在key
		if !panelViper.IsSet(key) {
			// 不存在
			return fmt.Errorf("写入配置失败，因为没有%v字段", key)
		}

		// 更新配置
		panelViper.Set(key, GetType(value))
	}

	// 写入配置
	panelViper.WriteConfig()

	return nil
}

// 获取interface的类型  然后转成对应类型后 返回
func GetType(value interface{}) any {
	// 类型断言
	switch v := value.(type) {
	case int:
		return v
	case string:
		return v
	case float64:
		if v == math.Trunc(v) {
			return int(v)
		}
		return v
	default:
		return v
	}
}

// 获取外部ip
func getExternalIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// 设置 web页面的请求接口
func setWebPort(ip string, port int) {

	// 打开文件执行覆盖操作
	file, err := os.OpenFile("./static/web/cfg.js", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入文件
	file.WriteString(fmt.Sprintf("const serverUrl = '%v:%v'", ip, port))
}
