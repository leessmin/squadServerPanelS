package init

import (
	"SSPS/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.uber.org/zap/zapcore"
)

// 服务器的状态  记录服务器 开启或关闭
func init() {
	logger.CreateLogger().Log(zapcore.InfoLevel, "面板服务器启动成功")

	var stopLock sync.Mutex
	stop := false
	stopChan := make(chan struct{}, 1)
	signalChan := make(chan os.Signal, 1)

	go func(stop *bool, stopChan chan struct{}, signalChan chan os.Signal) {
		//阻塞程序运行，直到收到终止的信号
		<-signalChan
		stopLock.Lock()
		*stop = true
		stopLock.Unlock()

		logger.CreateLogger().Log(zapcore.InfoLevel, "面板服务器关闭")
		// 刷新日志缓存
		logger.CreateLogger().Sync()

		stopChan <- struct{}{}
		os.Exit(0)
	}(&stop, stopChan, signalChan)

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
}
