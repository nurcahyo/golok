package manager

import "github.com/nurcahyo/golok/contract"

var (
	caches = make(map[string]contract.Loggable)
)

func AddCache(name string, logger contract.Loggable) {
	caches[name] = logger
}

func GetCache(name string) contract.Loggable {
	if val, ok := caches[name]; ok {
		return val
	}
	return nil
}
