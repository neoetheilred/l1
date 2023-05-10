package main

import "fmt"

func main() {
	in := make(chan int)         // Channel to send numbers to
	out := make(chan int)        // Channels to read proccessed numbers from
	nums := []int{1, 2, 3, 4, 5} // Source data

	go writer(nums, in)  // Write number to `in` channel
	go proccess(in, out) // Apply (*2) to nums from `in` channel, write results to `out` channel
	printer(out)         // Print numbers from `out` channel
}

// Writes numbers from `nums` to channel `ch`
func writer(nums []int, ch chan int) {
	for _, x := range nums {
		ch <- x
	}
	close(ch) // Close channel to avoid deadlock
}

// Receives numbers (x) from `input`, sends (x*2) to `output`
func proccess(input <-chan int, output chan<- int) {
	for x := range input {
		output <- x * 2
	}
	close(output) // Close channel to avoid deadlock
}

// Receives numbers from `ch` and prints them to stdout
func printer(ch <-chan int) {
	for x := range ch {
		fmt.Println(x)
	}
}
