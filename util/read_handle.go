package util

import (
	"bufio"
	"fmt"
	"os"
)

// 读取游戏服务器的配置文件

type ReadHandle struct {
}

var readHandle *ReadHandle

// 单例模式  创建一个  读取游戏服务器的配置文件  结构体实例
func CreateReadHandle() *ReadHandle {
	once.Do(func() {
		readHandle = &ReadHandle{}
	})

	return readHandle
}

// 读取配置文件  传入chan 将读取到的行通过chan 返回
func (rh *ReadHandle) ReadConfig(path string, ch chan string) {
	go func(path string, ch chan string) {
		// 打开文件
		file, err := os.Open(path)
		if err != nil {
			panic(fmt.Sprintf("打开文件失败,err:%v", err))
		}
		defer file.Close()

		buf := bufio.NewScanner(file)

		for {
			// 扫描到末尾，结束扫描
			if !buf.Scan() {
				// 关闭通道
				close(ch)
				break
			}

			// 获取当前行文字
			line := buf.Text()
			// fmt.Println(line)
			// 向通道发送读取到的文字
			ch <- line
		}
	}(path, ch)
}
