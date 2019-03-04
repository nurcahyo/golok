package lok

import (
	"testing"
)

func TestLog(t *testing.T) {
	config := map[string]interface{}{
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
				"level":    "error",
				"daily":    false,
				"filename": "log",
				"path":     "./",
				"driver":   "file",
			},
			// "sentry": map[string]interface{}{
			// 	"level":  "error",
			// 	"driver": "sentry",
			// },
		},
	}
	Initialize(config)
	Debugf("test %s", "wew")
}
