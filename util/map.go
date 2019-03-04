package util

func MapGet(obj map[string]interface{}, key string, defaultVal interface{}) interface{} {
	if val, ok := obj[key]; ok {
		return val
	}
	return defaultVal
}
