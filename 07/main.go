/*
 * Реализовать конкурентную запись данных в map.
 */
package main

import (
	"sync"
)

// Структура для удобной конкурентной работы с мапой
type syncMap struct {
	m  map[string]int
	mu *sync.RWMutex // Для неполного блокирования доступа при разных операциях
}

func (sm syncMap) Add(k string, v int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.m[k] = v
}

func (sm syncMap) Get(k string) int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	return sm.m[k]
}

func New() syncMap {
	return syncMap{
		m:  make(map[string]int),
		mu: &sync.RWMutex{},
	}
}

func main() {
}
