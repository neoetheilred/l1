package main

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	calcSum1(nums)
	calcSum2(nums)
}

func calcSum1(nums []int) {
	mx := sync.Mutex{}       // Use mutex to safely assign to res
	wg := sync.WaitGroup{}   // Use waitgroup to wait until all goroutines are done
	res := 0                 // Result will be stored here
	for _, n := range nums { // Run goroutine for each number
		wg.Add(1) // Add one more goroutine to wait for
		go func(k int) {
			defer wg.Done() // After calculations notify waitgroup that the goroutine is done
			mx.Lock()       // Lock mutex to safely access res
			res += k * k
			mx.Unlock() // Unlock mutex
		}(n)
	}
	wg.Wait() // Wait until all goroutines are done
	fmt.Println(res)
}

func calcSum2(nums []int) {
	ch := make(chan int, len(nums)) // Make buffered channel, which will accept numbers
	dest := make(chan int)          // Channel for result squared numbers
	res := 0                        // Result will be stored here
	//Write numbers to channels
	for _, n := range nums {
		ch <- n
	}
	// Run goroutines (any amount of them)
	go calcSquaresFromChannel(ch, dest)
	go calcSquaresFromChannel(ch, dest)
	go calcSquaresFromChannel(ch, dest)

	// Read the squares from dest, accumulate result
	for i := 0; i < len(nums); i++ {
		res += <-dest
	}

	fmt.Println(res)
}

// Accepts numbers from ch, writes squared numbers to dest
func calcSquaresFromChannel(ch <-chan int, dest chan<- int) {
	for n := range ch {
		dest <- n * n
	}
}
