package syslog

import (
	"fmt"
	"log"

	"github.com/nurcahyo/golok/util"
)

type SyslogHandler struct {
	level string
}

func format(level string, message string) string {
	return fmt.Sprintf("[%s] %s", level, message)
}

func write(level string, message string) {
	log.Println(format(level, message))
}

// NewHandler Create a new Stack Handler
func NewHandler(config map[string]interface{}) *SyslogHandler {
	return &SyslogHandler{
		level: util.MapGet(config, "level", "error").(string),
	}
}

func (handler *SyslogHandler) Debug(message string) {
	if util.LevelGte(handler.level, "debug") {
		write("debug", message)
	}
}

func (handler *SyslogHandler) Info(message string) {
	if util.LevelGte(handler.level, "info") {
		write("info", message)
	}
}

func (handler *SyslogHandler) Warning(message string) {
	if util.LevelGte(handler.level, "warning") {
		write("warning", message)
	}
}

func (handler *SyslogHandler) Error(message string) {
	if util.LevelGte(handler.level, "error") {
		write("error", message)
	}
}

func (handler *SyslogHandler) Critical(message string) {
	if util.LevelGte(handler.level, "critical") {
		write("critical", message)
	}
}

func (handler *SyslogHandler) Debugf(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "debug") {
		write("debug", fmt.Sprintf(format, params...))
	}
}

func (handler *SyslogHandler) Infof(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "info") {
		write("info", fmt.Sprintf(format, params...))
	}
}

func (handler *SyslogHandler) Warningf(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "warning") {
		write("warning", fmt.Sprintf(format, params...))
	}
}

func (handler *SyslogHandler) Errorf(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "error") {
		write("error", fmt.Sprintf(format, params...))
	}
}

func (handler *SyslogHandler) Criticalf(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "critical") {
		write("critical", fmt.Sprintf(format, params...))
	}
}
