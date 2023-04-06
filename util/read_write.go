package util

import (
	"SSPS/config"
	"bufio"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

// 读取游戏服务器的配置文件

type ReadWrite struct {
}

var readWrite *ReadWrite

// 单例模式  创建一个  读取游戏服务器的配置文件  结构体实例
func CreateReadWrite() *ReadWrite {
	once.Do(func() {
		readWrite = &ReadWrite{}
	})

	return readWrite
}

// 读取配置文件  传入chan 将读取到的行 通过chan 返回
func (rw *ReadWrite) ReadConfig(fileName string, ch chan string) {
	go func(fileName string, ch chan string) {
		// 路径拼接
		fileName = rw.basePathJoin(fileName)

		// 打开文件
		file, err := os.Open(fileName)
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

			// 判断是否为注释掉的行
			if isAnnotation(line) {
				// 跳过该行
				continue
			}

			// 向通道发送读取到的文字
			ch <- line
		}
	}(fileName, ch)
}

// 向配置文件追加一行
// fileName 文件名  index追加的行   content添加的内容
func (rw *ReadWrite) InsertLineConfig(fileName string, index int, content string) {
	// 路径拼接
	fileName = rw.basePathJoin(fileName)

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	buf := bufio.NewScanner(file)

	// 储存每一行的文本
	var lineArr []string
	// 读取的行 索引
	i := 1
	for buf.Scan() {

		// 获取当前行文字
		line := buf.Text()

		lineArr = append(lineArr, line)

		// 判断是否到达追加行索引
		if i == index {
			lineArr = append(lineArr, content)
		}

		i++
	}

	// 读取完毕 关闭文件
	file.Close()

	// 重新打开文件执行覆盖操作
	file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入文件
	file.WriteString(strings.Join(lineArr, "\n"))
}

// 正则表达式
// 匹配： Group=Admin:kick,ban,changemap  // 管理员
// ^Group=Admin:([A-z]+,{0,}){0,}([^\n]*\/\/[^\n]*){0,}

// 传入正则表达式  pattern 正则表达式
// 查找符合正则表达式内容的行数
// 多个内容符合 则 返回最后一个内容的索引
// 返回 -1 代表不存在内容
func (rw *ReadWrite) FindContentIndex(pattern, fileName string) int {
	// 路径拼接
	fileName = rw.basePathJoin(fileName)

	// 打开文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("打开文件失败,err:%v", err))
	}
	defer file.Close()

	buf := bufio.NewScanner(file)

	// 内容的 行数
	index := 1
	// 存在的内容的行数 多个
	var indexArr []int

	// 逐行扫描文件
	for buf.Scan() {

		// 获取当前行文字
		line := buf.Text()

		isOk, err := regexp.MatchString(pattern, line)
		if err != nil {
			panic(fmt.Sprint("使用正则表达式出错,err:", err))
		}

		// 判断是否符合正则表达式的内容
		if isOk {
			// 记录索引
			indexArr = append(indexArr, index)
		}

		// 索引++
		index++
	}

	// 判断是否有内容
	if len(indexArr) <= 0 {
		return -1
	}

	return indexArr[len(indexArr)-1]
}

// 拼接路径  传入需要找到的文件，，底层路径为配置文件中的game_serve_path
func (rw *ReadWrite) basePathJoin(fileName string) string {
	// 默认配置文件路径 .游戏根目录 ./SquadGame
	p := path.Join(config.PanelConf.GameServePath, "SquadGame", fileName)
	return p
}

// 判断是否为注释字符串 或者 空行   如果是则返回true
func isAnnotation(str string) bool {

	// 去除字符串两端空格
	str = strings.TrimSpace(str)

	// 判断是否为空行
	if len(str) <= 0 {
		return true
	}

	// 判断前两个字符是否为 //
	isFind := strings.Contains(str[0:2], "//")
	// 判断前两个字符是否包含 #
	isFind01 := strings.Contains(str[0:2], "#")

	return isFind || isFind01
}
