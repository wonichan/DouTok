package logger

import "github.com/cloudwego/hertz/pkg/common/hlog"

var logger hlog.FormatLogger

func InitLogger() {
	logger = hlog.DefaultLogger()
}

func GetLogger() hlog.FormatLogger {
	return logger
}
