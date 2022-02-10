package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

type LogSetting struct {
	// Path log file output location
	Path string
	// Level log output level
	Level string
	// Formatter test or json
	Formatter string
	// Output os or file
	Output string
}

func SetLog(path, level, formatter, output string) {

	// 判断日志级别
	switch level {
	case "panic":
		logrus.SetLevel(logrus.PanicLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetLevel(logrus.TraceLevel)
	}

	// 设置日志输出位置
	if output == "os" {
		logrus.SetOutput(os.Stdout)
	} else if output == "file" {

		file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			logrus.Errorf("log.logSet failed, err: %v", err)
		}

		logrus.SetOutput(file)
	}

	// 设置日志显示格式
	switch formatter {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "txt":
		logrus.SetFormatter(&logrus.TextFormatter{})

	}
}
