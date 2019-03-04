package manager

import (
	"fmt"

	"github.com/nurcahyo/golok/contract"
	"github.com/nurcahyo/golok/manager/factory"
)

type logManager struct {
	config map[string]interface{}
}

func NewLogManager(cfg map[string]interface{}) *logManager {
	return &logManager{
		config: cfg,
	}
}

func (m *logManager) GetConfig(key string, defaultVal string) interface{} {
	if val, ok := m.config[key]; ok {
		return val
	}
	return defaultVal
}

func (m *logManager) GetLog(logName string) contract.Loggable {
	if GetCache(logName) == nil {
		AddCache(logName, m.makeLog(logName))
	}
	return GetCache(logName)
}

func (m *logManager) makeLog(logName string) contract.Loggable {
	if _, ok := m.config["channels"]; !ok {
		panic("Can't find channels configuration, channels configuration is empty")
	}
	channels := m.config["channels"].(map[string]interface{})
	if _, ok := channels[logName]; !ok {
		panic(fmt.Sprintf("Can't find channels configuration for log handler %s", logName))
	}
	channelConfig := channels[logName].(map[string]interface{})
	driver, ok := channelConfig["driver"].(string)
	if !ok {
		panic(fmt.Sprintf("Can't driver configuration on channel %s configuration.", logName))
	}

	return factory.MakeHandler(driver, channelConfig, m)
}
