package log

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"sync"
	"time"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

func Init(cfg Config) *logrus.Logger {
	once.Do(func() {
		// 设置日志格式为json格式
		logger = logrus.New()
		logger.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			TimestampFormat: "2006-01-02 15:04:05",
			CallerFirst:     true,
		})
		logger.SetReportCaller(true)
		logger.SetOutput(os.Stdout)
		logger.SetLevel(6)

		writer, _ := rotatelogs.New(
			cfg.Path+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(cfg.Path),
			rotatelogs.WithMaxAge(time.Hour*time.Duration(cfg.MaxAge)),
			rotatelogs.WithRotationTime(time.Hour*time.Duration(cfg.RotationTime)),
		)

		fileAndStdoutWriter := io.MultiWriter(os.Stdout, writer)
		logrus.SetOutput(fileAndStdoutWriter)
	})
	return logger
}
