package log

type Config struct {
	Path         string
	MaxAge       int64
	RotationTime int64
}
