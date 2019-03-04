package syslog

import (
	"log"
	"os"

	"github.com/nurcahyo/golok/util"
)

type SyslogHandler struct {
	level  string
	logger *log.Logger
}

func (handler *SyslogHandler) write(level string, message string) {
	if util.LevelLte(handler.level, level) {
		log.Output(5, util.Format(level, message))
	}
}

// NewHandler Create a new Stack Handler
func NewHandler(config map[string]interface{}) *SyslogHandler {
	return &SyslogHandler{
		level:  util.MapGet(config, "level", "error").(string),
		logger: log.New(os.Stderr, "", log.Llongfile|log.LUTC|log.Lmicroseconds),
	}
}

func (handler *SyslogHandler) Debug(message string) {
	handler.write("debug", message)
}

func (handler *SyslogHandler) Info(message string) {
	handler.write("info", message)
}

func (handler *SyslogHandler) Warning(message string) {
	handler.write("warning", message)
}

func (handler *SyslogHandler) Error(err error) {
	handler.write("error", err.Error())
}

func (handler *SyslogHandler) Critical(err error) {
	handler.write("critical", err.Error())
}
