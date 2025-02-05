package logger

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestInfo(t *testing.T) {
	Info("test")
}

func TestDebug(t *testing.T) {
	InitWithConfig(Config{
		Debug: true,
		Level: int8(zapcore.DebugLevel),
	})
	Debug("test")
}
