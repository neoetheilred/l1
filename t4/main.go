package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: <executable name> N")
	}
	n, err := strconv.Atoi(os.Args[1]) // Get amount of workers from command-line args
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan interface{}) // Channel with data
	// Run `n` workers
	for i := 0; i < n; i++ {
		go worker(ch, i+1)
	}
	generateData(ch) // Generate data
}

// Infinitely writes data to channel
// When user presses ^C exits
func generateData(ch chan<- interface{}) {
	intChan := make(chan os.Signal, 1)
	signal.Notify(intChan, os.Interrupt)
	for {
		select {
		case <-intChan:
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Reads data from channel and prints to stdout
func worker(ch <-chan interface{}, order int) {
	for c := range ch {
		fmt.Printf("[worker %d]: %v\n", order, c)
	}
}
