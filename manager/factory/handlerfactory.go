package factory

import (
	"fmt"

	"github.com/nurcahyo/golok/contract"
	"github.com/nurcahyo/golok/handlers/stack"
	"github.com/nurcahyo/golok/handlers/syslog"
	"github.com/nurcahyo/golok/handlers/file"
	"github.com/nurcahyo/golok/handlers/sentry"
)

func MakeHandler(driver string, channelConfig map[string]interface{}, manager contract.Manager) contract.Loggable {
	switch driver {
	case "stack":
		return stack.NewHandler(channelConfig["channels"].([]string), manager)
	case "syslog":
		return syslog.NewHandler(channelConfig)
	case "file":
		return file.NewHandler(channelConfig)
	case "sentry":
		return sentry.NewHandler(channelConfig)
	}
	panic(fmt.Sprintf("Handler for %s driver not found.", driver))
}
