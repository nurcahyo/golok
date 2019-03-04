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

func (handler *StackHandler) Info(message string) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Info(message)
	}
}

func (handler *StackHandler) Warning(message string) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Warning(message)
	}
}

func (handler *StackHandler) Error(err error) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Error(err)
	}
}

func (handler *StackHandler) Critical(err error) {
	for _, channel := range handler.channels {
		handler.manager.GetLog(channel).Critical(err)
	}
}
