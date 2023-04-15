package logger

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestLog(t *testing.T) {
	CreateLogger().Log(zapcore.InfoLevel, "测试")
}
