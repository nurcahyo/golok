package file

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/nurcahyo/golok/util"
)

type FileHandler struct {
	level  string
	logger *log.Logger
}

func format(level string, message string) string {
	return fmt.Sprintf("[%s] %s", level, message)
}

func (handler FileHandler) write(level string, message string) {
	handler.logger.Println(format(level, message))
}

func getWritter(path string) io.Writer {
	if !util.FileExists(path) {
		util.CreateFile(path)
	}
	file, err := util.OpenFile(path)
	if err != nil {
		log.Fatalf("Error opening log file %v", err)
	}
	return file
}

// NewHandler Create a new Stack Handler
func NewHandler(config map[string]interface{}) *FileHandler {
	if _, ok := config["path"]; !ok {
		log.Fatal("File log configuration doesn't have correct path value.")
	}
	isDaily := util.MapGet(config, "isDaily", false).(bool)
	fileName := util.MapGet(config, "filename", "log").(string)
	if isDaily {
		fileName = fmt.Sprintf("%s-%s", time.Now().Format("2006-01-02"))
	}
	filePath := fmt.Sprintf("%s/%s.log", config["path"].(string), fileName)
	logger := log.New(getWritter(filePath), "", log.Lshortfile|log.LUTC)
	return &FileHandler{
		level:  util.MapGet(config, "level", "error").(string),
		logger: logger,
	}
}

func (handler *FileHandler) Debug(message string) {
	if util.LevelGte(handler.level, "debug") {
		handler.write("debug", message)
	}
}

func (handler *FileHandler) Info(message string) {
	if util.LevelGte(handler.level, "info") {
		handler.write("info", message)
	}
}

func (handler *FileHandler) Warning(message string) {
	if util.LevelGte(handler.level, "warning") {
		handler.write("warning", message)
	}
}

func (handler *FileHandler) Error(message string) {
	if util.LevelGte(handler.level, "error") {
		handler.write("error", message)
	}
}

func (handler *FileHandler) Critical(message string) {
	if util.LevelGte(handler.level, "critical") {
		handler.write("critical", message)
	}
}

func (handler *FileHandler) Debugf(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "debug") {
		handler.write("debug", fmt.Sprintf(format, params...))
	}
}

func (handler *FileHandler) Infof(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "info") {
		handler.write("info", fmt.Sprintf(format, params...))
	}
}

func (handler *FileHandler) Warningf(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "warning") {
		handler.write("warning", fmt.Sprintf(format, params...))
	}
}

func (handler *FileHandler) Errorf(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "error") {
		handler.write("error", fmt.Sprintf(format, params...))
	}
}

func (handler *FileHandler) Criticalf(format string, params ...interface{}) {
	if util.LevelGte(handler.level, "critical") {
		handler.write("critical", fmt.Sprintf(format, params...))
	}
}
