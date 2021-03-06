package golok

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

func ensureLogManager() {
	if logManager != nil {
		return
	}

	logManager = manager.NewLogManager(map[string]interface{}{
		"default": "system",
		"channels": map[string]interface{}{
			"system": map[string]interface{}{
				"level":  "debug",
				"driver": "syslog",
			},
		},
	})
}

func defaultLog() contract.Loggable {
	if defaultLogger == nil {
		ensureLogManager()
		defaultHandlerName := logManager.GetConfig("default", "syslog").(string)
		defaultLogger = logManager.GetLog(defaultHandlerName)
	}
	return defaultLogger
}

func Debug(msg string) {
	defaultLog().Debug(msg)
}

func Error(err error) {
	defaultLog().Error(err)
}

func Critical(err error) {
	defaultLog().Critical(err)
}

func Warning(msg string) {
	defaultLog().Warning(msg)
}

func Info(msg string) {
	defaultLog().Info(msg)
}

func Stack(channels []string) *stack.StackHandler {
	return stack.NewHandler(channels, logManager)
}
