package logger

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志

// 日志结构体
type loggerStruct struct {
	logger *zap.Logger
}

var (
	logger  *loggerStruct
	logOnce sync.Once
)

// 创建日志
func CreateLogger() *loggerStruct {
	logOnce.Do(func() {
		encoderConfig := zap.NewProductionEncoderConfig()
		// 设置日志的时间格式
		encoderConfig.TimeKey = "time"
		// 设置日志记录中时间的格式
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
		// 日志Encoder 还是JSONEncoder，把日志行格式化成JSON格式的
		encoder := zapcore.NewJSONEncoder(encoderConfig)

		// 写入功能
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   "./panel_log/log.log",
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     30,   // days
			LocalTime:  true, // 本地时区
		})

		// 日志写入核心
		writeCore := zapcore.NewCore(encoder, w, zap.InfoLevel)

		// 创建日志核心
		core := zapcore.NewTee(writeCore)

		// 创建日志库对象
		logger = &loggerStruct{
			logger: zap.New(core),
		}
	})

	return logger
}

// 记录 日志
func (l *loggerStruct) Log(level zapcore.Level, msg string) {
	l.logger.Log(level, msg)
	l.logger.Sync()
}

// 刷新日志缓存
func (l *loggerStruct) Sync() {
	l.logger.Sync()
}

// 日志信息
type logInfo struct {
	// 级别
	Level string `json:"level"`
	// 时间
	Time string `json:"time"`
	// 消息
	Msg string `json:"msg"`
}

// 读取日志
func (l *loggerStruct) ReadLog() []logInfo {
	// 路径拼接
	filePath := "./panel_log/log.log"

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		// 打开文件失败
		panic(fmt.Sprintf("打开文件失败,err:%v", err))
	}
	defer file.Close()

	buf := bufio.NewScanner(file)

	// 储存 日志
	var logArr []logInfo

	for {
		// 扫描到末尾，结束扫描
		if !buf.Scan() {
			break
		}

		// 获取当前行文字
		line := buf.Bytes()

		var info logInfo
		err := json.Unmarshal(line, &info)
		if err != nil {
			CreateLogger().Log(zap.PanicLevel, "日志解析失败")
			continue
		}

		// 追加
		logArr = append(logArr, info)
	}

	return logArr
}
