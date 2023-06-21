package util

import (
	"SSPS/config"
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"sync"
)

// 读取游戏服务器的配置文件

// 读写 结构体
type ReadWrite struct {
}

var (
	readWrite *ReadWrite
	onceRW    sync.Once
)

// 单例模式  创建一个  读取游戏服务器的配置文件  结构体实例
func CreateReadWrite() *ReadWrite {
	onceRW.Do(func() {
		readWrite = &ReadWrite{}
	})

	return readWrite
}

// TODO: 已知错误 当读取不到文件时 会使服务器宕机
// 读取配置文件  传入chan 将读取到的行 通过chan 返回
func (rw *ReadWrite) ReadConfig(fileName string, ch chan string) {
	go func(fileName string, ch chan string) {
		// 捕获 panic
		defer func(ch chan string) {
			if err := recover(); err != nil {
				fmt.Println(err)
				close(ch)
			}
		}(ch)

		// 路径拼接
		fileName = rw.basePathJoin(fileName)

		// 打开文件
		file, err := os.Open(fileName)
		if err != nil {
			// 打开文件失败
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

			// 向通道发送读取到的文字
			ch <- line
		}
	}(fileName, ch)
}

// 读取配置文件，不读取注释部分 传入chan 将读取到的行 通过chan 返回    （对ReadConfig进行封装）
func (rw *ReadWrite) ReadNotCommentConfig(fileName string, ch chan string) {
	go func(fileName string, ch chan string) {
		ch0 := make(chan string)

		rw.ReadConfig(fileName, ch0)

		for {
			// 获取数据
			line, ok := <-ch0

			// 通道关闭 跳出for循环
			if !ok {
				// 关闭通道
				close(ch)
				break
			}

			// 判断是否为注释掉的行
			if IsAnnotation(line) {
				// 跳过该行
				continue
			}

			// 读取到数据  发送到ch通道
			ch <- line
		}
	}(fileName, ch)
}

// 传入需要拷贝的配置文件名字   备份到的路径
// 拷贝文件
func (rw *ReadWrite) CopyFile(srcFile, destPath string) error {
	// 路径拼接
	srcPath := rw.basePathJoin(srcFile)

	sourceFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}

	destFile := path.Join(destPath, srcFile)

	file, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	defer sourceFile.Close()
	defer file.Close()

	io.Copy(file, sourceFile)

	return nil
}

// 向配置文件 追加一行  或 替换一行
// fileName 文件名  index追加的行   content添加的内容   action执行动作 决定该动作是写入还是替换  InsertReplaceHandle
func (rw *ReadWrite) InsertReplaceLineConfig(fileName string, index int, content string, action InsertReplaceHandle) {

	ch := make(chan string)

	// 读取文件
	rw.ReadConfig(fileName, ch)

	// 处理 获取到替换的内容
	str := action.Handle(index, content, &ch)

	// 写入文件
	rw.coverWrite(fileName, str)
}

// 替换内容的方法 接口
type InsertReplaceHandle interface {
	//     索引 内容     读取到文本的通道
	Handle(int, string, *chan string) string
}

// 插入某一行到文件
type InsertLine struct{}

func (il *InsertLine) Handle(index int, content string, ch *chan string) string {
	// 储存每一行的文本
	var lineArr []string
	// 读取的行 索引
	i := 1
	for {
		// 获取数据
		line, ok := <-*ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		// 插入
		lineArr = append(lineArr, line)

		// 判断是否到达追加行索引
		if i == index {
			lineArr = append(lineArr, content)
		}

		i++
	}

	return strings.Join(lineArr, "\n")
}

// 替换某一行到文件
type ReplaceLine struct{}

func (rl *ReplaceLine) Handle(index int, content string, ch *chan string) string {
	// 储存每一行的文本
	var lineArr []string
	// 读取的行 索引
	i := 1
	for {
		// 获取数据
		line, ok := <-*ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		// 判断是否到达替换行索引
		if i == index {
			// 替换成 content
			lineArr = append(lineArr, content)
		} else {
			lineArr = append(lineArr, line)
		}

		i++
	}

	return strings.Join(lineArr, "\n")
}

// 替换文件所有内容
type ReplaceAll struct{}

func (ra ReplaceAll) Handle(index int, content string, ch *chan string) string {
	// 读取的行 索引

	// 读取端 无法关闭通道，就让通道读取完自动关闭
	go func(ch *chan string) {
		for {
			_, ok := <-*ch
			// 通道关闭 跳出for循环
			if !ok {
				break
			}
		}
	}(ch)

	return content
}

// 追加新一行
type AppendLine struct{}

func (al *AppendLine) Handle(index int, content string, ch *chan string) string {
	// 储存每一行的文本
	var lineArr []string
	// 读取的行 索引
	i := 1
	for {
		// 获取数据
		line, ok := <-*ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		lineArr = append(lineArr, line)

		i++
	}

	// 向文本末尾处追加一行
	lineArr = append(lineArr, content)

	return strings.Join(lineArr, "\n")
}

// 删除 文件中的 一行
type DeleteLine struct{}

func (dl *DeleteLine) Handle(index int, content string, ch *chan string) string {
	// 储存每一行的文本
	var lineArr []string
	// 读取的行 索引
	i := 1
	for ; ; i++ {
		// 获取数据
		line, ok := <-*ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		// 判断 是否为删除的 索引
		// 是则不储存当前行的文本
		if index == i {
			continue
		}
		lineArr = append(lineArr, line)

	}

	return strings.Join(lineArr, "\n")
}

// 删除 符合 正则表达式的内容
type DeleteRegular struct{}

func (dr *DeleteRegular) Handle(index int, pattern string, ch *chan string) string {
	// 储存每一行的文本
	var lineArr []string
	// 读取的行 索引
	i := 1
	for ; ; i++ {
		// 获取数据
		line, ok := <-*ch

		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		isOk := CreateRegexp().VerifyStr(pattern, line)

		// 判断是否符合正则表达式的内容
		if isOk {
			// 符合正则表达式 删除
			// 跳过这次for循环
			continue
		}

		lineArr = append(lineArr, line)

	}

	return strings.Join(lineArr, "\n")
}

// 判断目录是否存在
func (rw *ReadWrite) IsDir(dirPath string) (bool, error) {
	s, err := os.Stat(dirPath)
	if err != nil {
		return false, err
	}

	return s.IsDir(), nil
}

// 创建目录
func (rw *ReadWrite) CreateDir(dirName string) error {
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		return err
	}

	return nil
}

// 正则表达式
// 匹配： Group=Admin:kick,ban,changemap  // 管理员
// ^Group=Admin:([A-z]+,{0,}){0,}([^\n]*\/\/[^\n]*){0,}

// 传入正则表达式  pattern 正则表达式
// 查找符合正则表达式内容的行数
// 多个内容符合 则 返回最后一个内容的索引
// 返回 -1 代表不存在内容
func (rw *ReadWrite) FindContentIndex(pattern, fileName string) int {

	// 内容的 行数
	index := 1
	// 记录 符合正则表达式的行 索引
	recordIndex := -1

	ch := make(chan string)

	// 读取文件
	rw.ReadConfig(fileName, ch)
	for {
		line, ok := <-ch
		// 通道关闭 跳出for循环
		if !ok {
			break
		}

		isOk := CreateRegexp().VerifyStr(pattern, line)

		// 判断是否符合正则表达式的内容
		if isOk {
			// 记录索引
			recordIndex = index
		}

		// 索引++
		index++
	}

	return recordIndex
}

// 拼接路径  传入需要找到的文件，，底层路径为配置文件中的game_serve_path
func (rw *ReadWrite) basePathJoin(fileName string) string {
	// TODO:已知问题，配置文件路径可能不对
	// 默认配置文件路径 .游戏根目录 ./SquadGame
	p := path.Join(config.CreatePanelConf().GameServePath, "SquadGame", "ServerConfig", fileName)
	return p
}

// 覆盖 文件
func (rw *ReadWrite) coverWrite(fileName, content string) {

	// 打开文件执行覆盖操作
	file, err := os.OpenFile(rw.basePathJoin(fileName), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入文件
	file.WriteString(content)
}

// 判断是否为注释字符串 或者 空行   如果是则返回true
func IsAnnotation(str string) bool {

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
