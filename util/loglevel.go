package util

var logLevelScore = map[string]int8{
	"debug":    0,
	"info":     1,
	"warning":  2,
	"error":    3,
	"critical": 4,
}

func LevelGte(a string, b string) bool {
	return a >= b
}
