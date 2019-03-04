package contract

type Loggable interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(err error)
	Critical(err error)
}

type Manager interface {
	GetLog(logName string) Loggable
	GetConfig(key string, defaultVal string) interface{}
}

type MutateSentryTagsMiddleware func(tag map[string]string)
