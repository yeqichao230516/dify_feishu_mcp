package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func NewLogger() {
	Logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		DataKey:         "data",
	})

	setLogLevel()

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Logger.Fatal("无法打开日志文件:", err)
	}
	Logger.SetOutput(file)
}

// 设置日志级别
func setLogLevel() {
	levelStr := os.Getenv("LOG_LEVEL")
	if levelStr == "" {
		Logger.SetLevel(logrus.InfoLevel)
		return
	}

	level, err := logrus.ParseLevel(levelStr)
	if err != nil {
		Logger.Warnf("无效的日志级别: %s，使用默认Info级别", levelStr)
		Logger.SetLevel(logrus.InfoLevel)
		return
	}

	Logger.SetLevel(level)
	Logger.Debugf("设置日志级别为: %s", level.String())
}
