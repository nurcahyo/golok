package stack

import (
	"github.com/nurcahyo/golok/contract"
)

type StackHandler struct {
	channels []string
	manager  contract.Manager
}

// NewHandler Create a new Stack Handler
func NewHandler(channels []string, manager contract.Manager) *StackHandler {
	return &StackHandler{
		channels: channels,
		manager:  manager,
	}
}

func (handler *StackHandler) Debug(message string) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Debug(message)
	}
}

func (handler *StackHandler) Error(message string) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Error(message)
	}
}

func (handler *StackHandler) Critical(message string) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Critical(message)
	}
}

func (handler *StackHandler) Warning(message string) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Warning(message)
	}
}

func (handler *StackHandler) Debugf(format string, params ...interface{}) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Debugf(format, params...)
	}
}

func (handler *StackHandler) Errorf(format string, params ...interface{}) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Errorf(format, params...)
	}
}

func (handler *StackHandler) Criticalf(format string, params ...interface{}) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Criticalf(format, params...)
	}
}

func (handler *StackHandler) Warningf(format string, params ...interface{}) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Warningf(format, params...)
	}
}

func (handler *StackHandler) Info(message string) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Info(message)
	}
}

func (handler *StackHandler) Infof(format string, params ...interface{}) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Infof(format, params...)
	}
}
