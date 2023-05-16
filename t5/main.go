package main

/*
	Разработать программу, которая будет последовательно отправлять значения в канал,
	а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.
*/

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan interface{})     // channel, used to interact between writer and reader
	interrupt := make(chan struct{}) // Channel used to stop goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	// Run goroutines
	go func() {
		writer(ch, interrupt)
		fmt.Println("Writer exits")
		wg.Done()
	}()
	go func() {
		reader(ch, interrupt)
		fmt.Println("Reader exits")
		wg.Done()
	}()
	time.Sleep(time.Duration(n) * time.Second) // Sleep N seconds
	close(interrupt)                           // When channel is closed, all goroutines are ended
	wg.Wait()
}

// Continiously writes data to channel
func writer(ch chan<- interface{}, interrupt <-chan struct{}) {
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
	for {
		select {
		case c := <-ch:
			fmt.Println(c)
		case <-interrupt:
			return
		}
	}
}
