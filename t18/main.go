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

func main() {
	c := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				c.Inc()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(c.counter)
}
