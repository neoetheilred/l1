package main

import (
	"fmt"
	"time"
)

func main() {
	bench(func() {
		sleep(1000)
	})

	bench(func() {
		sleepTimer(1000)
	})
}

func bench(f func()) {
	start := time.Now()
	f()
	fmt.Printf("Elapsed time: %v\n", time.Since(start))
}

// Using timer for pausing goroutine
func sleepTimer(milliseconds int) {
	<-time.NewTimer(time.Millisecond * time.Duration(milliseconds)).C
}

// Use infinite loop, on each iteration check if required time has passed
func sleep(milliseconds int) {
	now := time.Now()
	for {
		if time.Since(now) >= time.Millisecond*time.Duration(milliseconds) {
			return
		}
	}
}
