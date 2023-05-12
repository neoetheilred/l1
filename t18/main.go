package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
	mx      sync.RWMutex
}

func (c *Counter) Inc() {
	c.mx.Lock()
	c.counter++
	c.mx.Unlock()
}

func (c *Counter) Value() int {
	c.mx.RLock()
	defer c.mx.RUnlock()
	return c.counter
}

func main() {
	c := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				c.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait() // Waiting for all goroutines to finish
	fmt.Println(c.Value())
}
