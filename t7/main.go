package main

import (
	"fmt"
	"sync"
)

type SyncMap[T comparable, U any] struct {
	m     map[T]U      // Underlying map
	mutex sync.RWMutex // Using RWMutex instead of Mutex to allow simultaneous access for read
}

func NewSyncMap[T comparable, U any]() *SyncMap[T, U] {
	return &SyncMap[T, U]{
		m: make(map[T]U),
	}
}

func (sm *SyncMap[T, U]) Set(key T, value U) {
	sm.mutex.Lock() // Lock for write
	sm.m[key] = value
	sm.mutex.Unlock() // Unlock for write
}

func (sm *SyncMap[T, U]) Get(key T) U {
	sm.mutex.RLock()         // Lock for read
	defer sm.mutex.RUnlock() // Unlock for read after reading operation is done
	return sm.m[key]
}

func main() {
	m := NewSyncMap[int, int]()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			m.Set(n, n+1)
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 10; i++ {
		fmt.Println(m.Get(i))
	}
}
