package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan interface{})        // channel, used to interact between writer and reader
	interrupt := make(chan struct{}, 2) // We need to stop 2 goroutines, so we use chan with buffer of 2
	// Run goroutines
	go writer(ch, interrupt)
	go reader(ch, interrupt)
	time.Sleep(time.Duration(n) * time.Second)
	interrupt <- struct{}{}
	interrupt <- struct{}{}
	time.Sleep(100 * time.Millisecond)
}

// Continiously writes data to channel
func writer(ch chan<- interface{}, interrupt <-chan struct{}) {
	defer fmt.Println("Writer exits")
	for {
		select {
		case <-interrupt: // Stop execution when smth is sent to interrupt channel
			return
		default: // Write data to channel
			ch <- map[string]string{"a": "b", "c": "d"} // Write data
			time.Sleep(time.Millisecond * 100)          // Sleep 100ms
		}
	}
}

// Reads from channel, interrupts when interrupt channel receives data
func reader(ch <-chan interface{}, interrupt <-chan struct{}) {
	defer fmt.Println("Reader exits")
	for {
		select {
		case c := <-ch:
			fmt.Println(c)
		case <-interrupt:
			return
		}
	}
}
