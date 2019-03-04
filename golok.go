package lok

import (
	"github.com/nurcahyo/golok/contract"
	"github.com/nurcahyo/golok/handlers/stack"
	"github.com/nurcahyo/golok/manager"
)

var (
	defaultLogger contract.Loggable
	logManager    contract.Manager
)

func Initialize(cfg map[string]interface{}) {
	logManager = manager.NewLogManager(cfg)
}

func defaultLog() contract.Loggable {
	if defaultLogger == nil {
		defaultHandlerName := logManager.GetConfig("default", "syslog").(string)
		defaultLogger = logManager.GetLog(defaultHandlerName)
	}

	return defaultLogger
}

func Debug(msg string) {
	defaultLog().Debug(msg)
}

func Error(msg string) {
	defaultLog().Error(msg)
}

func Critical(msg string) {
	defaultLog().Critical(msg)
}

func Warning(msg string) {
	defaultLog().Warning(msg)
}

func Info(msg string) {
	defaultLog().Info(msg)
}

func Debugf(format string, params ...interface{}) {
	defaultLog().Debugf(format, params...)
}
func Errorf(format string, params ...interface{}) {
	defaultLog().Errorf(format, params...)
}
func Criticalf(format string, params ...interface{}) {
	defaultLog().Criticalf(format, params...)
}
func Warningf(format string, params ...interface{}) {
	defaultLog().Warningf(format, params...)
}
func Infof(format string, params ...interface{}) {
	defaultLog().Infof(format, params...)
}

func Stack(channels []string) *stack.StackHandler {
	return stack.NewHandler(channels, nil)
}
