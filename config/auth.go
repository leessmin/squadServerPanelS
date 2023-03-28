package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// 权限 用户的信息
type AuthUser struct {
	Username string
	Password string
}

// auth 配置文件读取器
var authViper *viper.Viper

// 读取 auth 配置文件
func (a AuthUser) ReadAuthConfig() AuthUser {
	// 创建配置文件读取器
	authViper = viper.New()

	// 设置配置文件
	authViper.SetConfigName("auth")
	authViper.SetConfigType("toml")
	authViper.AddConfigPath("./panel_config/")

	// 读取配置文件
	err := authViper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// 将配置文件映射成user
	a.Username = authViper.GetString("account.username")
	a.Password = authViper.GetString("account.password")

	return a
}
