package util

import "fmt"

var logLevelScore = map[string]int8{
	"debug":    0,
	"info":     1,
	"warning":  2,
	"error":    3,
	"critical": 4,
}

func LevelLte(a string, b string) bool {
	return a <= b
}

func Format(level string, message string) string {
	return fmt.Sprintf("[%s] %s", level, message)
}
