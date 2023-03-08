package log

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

var (
	logger *logrus.Logger
	once   sync.Once
)

func Api(cfg *Config) *logrus.Logger {
	once.Do(func() {
		// 设置日志格式为json格式
		logger = logrus.New()
		logger.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			TimestampFormat: "[2006-01-02 15:04:05]",
			CallerFirst:     true,
			CustomCallerFormatter: func(f *runtime.Frame) string {
				return fmt.Sprintf("[%s:%d]", path.Base(f.File), f.Line)
			},
		})
		logger.SetReportCaller(true)
		logger.SetLevel(6)

		writer, _ := rotatelogs.New(
			cfg.Path+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(cfg.Path),
			rotatelogs.WithMaxAge(time.Second*time.Duration(cfg.MaxAge)),
			rotatelogs.WithRotationTime(time.Second*time.Duration(cfg.RotationTime)),
		)

		fileAndStdoutWriter := io.MultiWriter(os.Stdout, writer)
		logger.SetOutput(fileAndStdoutWriter)
	})
	return logger
}
