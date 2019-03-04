package lok

import (
	"errors"
	"testing"
)

var config = map[string]interface{}{
	"default": "stack",
	"channels": map[string]interface{}{
		"stack": map[string]interface{}{
			"channels": []string{"system", "file"},
			"driver":   "stack",
		},
		"system": map[string]interface{}{
			"level":  "debug",
			"driver": "syslog",
		},
		"file": map[string]interface{}{
			"level":    "debug",
			"daily":    false,
			"filename": "log",
			"path":     "./",
			"driver":   "file",
		},
		"sentry": map[string]interface{}{
			"level":       "error",
			"driver":      "sentry",
			"environment": "golok-test",
			"wait":        true,
			"dsn":         "https://772b29d6912f4efba93f82716ecfbb6f:e6c87cac79e448859c21ffe53ccee741@sentry.io/1407424",
		},
	},
}

func TestLog(t *testing.T) {
	Initialize(config)
	Debug("test log")
	Error(errors.New("Iam test log from local"))
}

func TestLogWithSentry(t *testing.T) {
	Initialize(config)
	Stack([]string{"sentry"}).Debug("Test stack with sentry log below config level")
	Stack([]string{"sentry"}).Error(errors.New("Test stack with sentry log success from local"))
}
