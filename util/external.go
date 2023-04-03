package util

import (
	"fmt"
	"io"
	"net/http"
)

// 获取外部ip
func GetExternalIP() string {
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
