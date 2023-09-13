package zop

// Logger 定义了 blog 项目的日志接口，该接口只包含了支持的日志记录方法（还缺点东西）
type Logger interface {
	Debugw(msg string, keyAndValues ...interface{})
	Infow(msg string, keyAndValues ...interface{})
	Warnw(msg string, keyAndValues ...interface{})
	Errorw(msg string, keyAndValues ...interface{})
	Panicw(msg string, keyAndValues ...interface{})
	Fatalw(msg string, keyAndValues ...interface{})
	Sync()
}