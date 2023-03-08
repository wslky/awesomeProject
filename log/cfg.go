package log

type Config struct {
	Path         string `json:"path" yaml:"path"`
	MaxAge       int64  `json:"max_age" yaml:"maxAge"`
	RotationTime int64  `json:"rotation_time" yaml:"rotationTime"`
}
