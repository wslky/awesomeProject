package log

import formatter "github.com/antonfisher/nested-logrus-formatter"

type Config struct {
	Name                  string               `json:"name" yaml:"name"`                                      //日志名称
	Path                  string               `json:"path" yaml:"path"`                                      // 日志路径
	MaxAge                int64                `json:"max_age" yaml:"maxAge"`                                 // 日志文件最多保存时间
	RotationTime          int64                `json:"rotation_time" yaml:"rotationTime"`                     //多久新建一个日志文件
	EnableSetReportCaller bool                 `json:"enable_set_report_caller" yaml:"enableSetReportCaller"` // 是否打印调用链
	Level                 int                  `json:"level" yaml:"level"`                                    //日志级别
	Formatter             *formatter.Formatter `json:"formatter" yaml:"formatter"`                            //格式设置
}
