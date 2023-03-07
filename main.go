package main

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var logger *logrus.Logger

func init() {
	// 设置日志格式为json格式
	logger = logrus.New()
	logger.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		TimestampFormat: "2006-01-02 15:04:05",
		FieldsOrder:     []string{"component", "category"},
		CallerFirst:     true,
	})
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(6)
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, logger.Writer())
	logrus.SetOutput(fileAndStdoutWriter)
}

func main() {
	logger.Info("hello, logrus....")
	logger.Info("hello, logrus1....")
	logger.Error("hello, logrus1....")
	logger.Warn("err")
}
