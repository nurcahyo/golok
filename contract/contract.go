package contract

type Loggable interface {
	Debug(msg string)
	Error(msg string)
	Critical(msg string)
	Warning(msg string)
	Info(msg string)
	Debugf(format string, params ...interface{})
	Errorf(format string, params ...interface{})
	Criticalf(format string, params ...interface{})
	Warningf(format string, params ...interface{})
	Infof(format string, params ...interface{})
}

type Manager interface {
	GetLog(logName string) Loggable
	GetConfig(key string, defaultVal string) interface{}
}
