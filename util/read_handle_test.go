package util

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	rh := CreateReadHandle()

	// 创建通道
	ch1 := make(chan string)

	rh.ReadConfig("Admins.cfg", ch1)

	for {
		data, ok := <-ch1
		// 通道关闭
		if !ok {
			return
		}

		// // 存在注释字符
		// isFind := strings.Contains(data, "//")
		// // 空行
		// isNil := len(strings.TrimSpace(data))

		// // 存在注释字符  或  空行
		// if isFind || isNil <= 0 {
		// 	continue
		// }
		fmt.Println(data)
	}
}
