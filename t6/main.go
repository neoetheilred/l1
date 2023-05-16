package main

/*Реализовать все возможные способы остановки выполнения горутины. */

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // We'll use this to wait until goroutine stops

	// First case:
	//		Stop goroutine when `quit` channel receives data
	quit := make(chan struct{}) // Channel is used to stop goroutine
	wg.Add(1)
	go worker1(quit, &wg)
	time.Sleep(1000 * time.Millisecond)
	quit <- struct{}{}
	wg.Wait()

	// Second case:
	//		Use channel with data, when channel is closed goroutine is stopped
	ch := make(chan interface{})
	wg.Add(1)
	go worker2(ch, &wg)
	for i := 0; i < 5; i++ { // Provide fake data load
		ch <- i + 1
	}
	close(ch) // Closing channel to stop goroutine
	wg.Wait()

	// Third case:
	// 		Use context
	ctx, cancel := context.WithCancel(context.Background())
	go worker3(ctx)
	time.Sleep(1000 * time.Millisecond)
	cancel()
}

// Waits until quit channel receives a signal to stop
func worker1(quit <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()                      // Unlock main goroutine
	defer fmt.Println("Worker1 is done") // Notify when done
	for {
		select {
		case <-quit: // Check if done
			// wg.Done()
			return // Stop
		default:
			fmt.Println("Worker1")             // Do some stuff
			time.Sleep(200 * time.Millisecond) // Imitate workload
		}
	}
}

// Reads from channel `ch`, when channel is closed the goroutine is stopped
func worker2(ch chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()                      // Unlock main goroutine
	defer fmt.Println("Worker2 is done") // Notify when done
	for c := range ch {
		fmt.Println(c) // Imitating some proccess
	}
}

// Does smth util context `ctx` is done
func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // If stop signal is received
			fmt.Println("Worker3 is done")
			return
		default:
			fmt.Println("Worker3")
			time.Sleep(200 * time.Millisecond) // Imitate some proccess
		}
	}
}
