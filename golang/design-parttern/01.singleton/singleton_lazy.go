package singleton

import (
	"sync"
)

var (
	lazySingleton *Singleton
	once          sync.Once
)

func GetLazyInstance() *Singleton {
	once.Do(func() {
		lazySingleton = &Singleton{}
	})
	return lazySingleton
}
