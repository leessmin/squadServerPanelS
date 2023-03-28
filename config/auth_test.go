package config

import (
	"fmt"
	"testing"
)

// 测试 读取 auth 配置文件
func TestReadAuthConfig(t *testing.T) {
	a := AuthUser{}
	user := a.ReadAuthConfig()

	fmt.Println(user)
}
