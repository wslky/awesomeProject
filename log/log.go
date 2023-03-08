package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

var (
	logMap sync.Map
)

func Api(cfg *Config) *logrus.Logger {

	value, ok := logMap.Load(cfg.Name)
	if ok {
		return value.(*logrus.Logger)
	}

	// 设置日志格式为json格式
	logger := logrus.New()
	logger.SetFormatter(cfg.Formatter)
	logger.SetReportCaller(cfg.EnableSetReportCaller)
	logger.SetLevel(logrus.Level(cfg.Level))

	writer, err := rotatelogs.New(
		cfg.Path+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(cfg.Path),
		rotatelogs.WithMaxAge(time.Second*time.Duration(cfg.MaxAge)),
		rotatelogs.WithRotationTime(time.Second*time.Duration(cfg.RotationTime)),
	)
	if err != nil {
		log.Fatal(err)
	}
	fileAndStdoutWriter := io.MultiWriter(os.Stdout, writer)
	logger.SetOutput(fileAndStdoutWriter)
	logMap.Store(cfg.Name, logger)

	return logger
}

/**
Example:
logger.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			TimestampFormat: "[2006-01-02 15:04:05]",
			CallerFirst:     true,
			CustomCallerFormatter: func(f *runtime.Frame) string {
				return fmt.Sprintf("[%s:%d]", path.Base(f.File), f.Line)
			},
		})
*/
