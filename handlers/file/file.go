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
	if util.LevelLte(handler.level, level) {
		handler.logger.Output(5, util.Format(level, message))
	}
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
		fileName = fmt.Sprintf("%s-%s", fileName, time.Now().Format("2006-01-02"))
	}
	filePath := fmt.Sprintf("%s/%s.log", config["path"].(string), fileName)
	logger := log.New(getWritter(filePath), "", log.Llongfile|log.LUTC)
	return &FileHandler{
		level:  util.MapGet(config, "level", "error").(string),
		logger: logger,
	}
}

func (handler *FileHandler) Debug(message string) {
	handler.write("debug", message)
}

func (handler *FileHandler) Info(message string) {
	handler.write("info", message)
}

func (handler *FileHandler) Warning(message string) {
	handler.write("warning", message)
}

func (handler *FileHandler) Error(err error) {
	handler.write("error", err.Error())
}

func (handler *FileHandler) Critical(err error) {
	handler.write("critical", err.Error())
}
