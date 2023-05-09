package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	// source slice
	nums := []int{2, 4, 6, 8, 10}
	fmt.Println(">>>1<<<")
	concurrentSquares1(nums) // Doesn't block main, using waitUserInterrupt() to wait until app is closed
	fmt.Println(">>>2<<<")
	concurrentSquares2(nums) // Blocks main until ends
	fmt.Println(">>>3<<<")
	concurrentSquares3(nums) // Blocks main until ends

	waitUserInterrupt() // Waits until user types ^C
}

// We will use buffered channels to process numbers
func concurrentSquares3(nums []int) {
	ch := make(chan int, len(nums)) // Buffered channel is used because we know the size of source data
	dest := make(chan int, len(nums))
	// Sending numbers to source channel
	for _, n := range nums {
		ch <- n
	}
	// Running goroutines (we can run any amount of goroutines there, since the code depends on channels)
	go chanSquares(ch, dest)
	go chanSquares(ch, dest)
	go chanSquares(ch, dest)
	go chanSquares(ch, dest)
	// We know the exact amount of results, so reading
	// only neccessary amount of times, avoiding deadlock
	for i := 0; i < len(nums); i++ {
		fmt.Println(<-dest)
	}
}

func chanSquares(ch <-chan int, dest chan<- int) {
	for n := range ch {
		dest <- n * n
	}
}

// Using waitgroups, similar to concurrentSquares1, but this version actually waits until goroutines end
func concurrentSquares2(nums []int) {
	wg := sync.WaitGroup{}
	wg.Add(2) // We will call 2 goroutines, so inc wg counter by 2
	go wgSquares(&wg, nums[:len(nums)/2])
	go wgSquares(&wg, nums[len(nums)/2:])
	wg.Wait() // Wait until all goroutines are done (counter of wg == 0)
}

func wgSquares(wg *sync.WaitGroup, nums []int) {
	defer wg.Done() // after calculations and output, notify waitgroup
	for _, n := range nums {
		fmt.Println(n * n)
	}
}

// Simple goroutines
// Running separate goroutines from two halves of slice
// Non-blocking
func concurrentSquares1(nums []int) {
	go printSquares(nums[:len(nums)/2])
	go printSquares(nums[len(nums)/2:])
}

func printSquares(a []int) {
	for _, n := range a {
		fmt.Println(n * n)
	}
}

func waitUserInterrupt() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch // blocks execution until os.Interrupt signal is received
}
