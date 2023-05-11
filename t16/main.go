package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 20; i++ {
		randomArr := generateRandomSlice(i)
		fmt.Printf("Source: %v\n", randomArr)
		quickSort(randomArr)
		fmt.Printf("Sorted: %v\n", randomArr)
		fmt.Printf("Is sorted: %v\n", isSorted(randomArr))
		fmt.Println()
	}
}

func generateRandomSlice(length int) []int {
	res := make([]int, length)
	for i := 0; i < len(res); i++ {
		res[i] = rand.Intn(100)
	}
	return res
}

func isSorted(arr []int) bool {
	if len(arr) == 0 {
		return true
	}
	p := arr[0]
	for i := 1; i < len(arr); i++ {
		if p > arr[i] {
			return false
		}
		p = arr[i]
	}
	return true
}

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	p := partition(arr)
	quickSort(arr[:p])
	quickSort(arr[p+1:])
}

func partition(arr []int) int {
	v := arr[len(arr)-1] // This element will be at the right place in source array after partition is over
	l := -1
	// After this loop we'll get smth like this:
	// [a1,a2,a3..., b1,b2,b3..., v], where a1..a_n < v, b1..b_n >= v
	for r := 0; r < len(arr); r++ {
		if arr[r] < v {
			l++
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
	// Place the element on the right place in source array, we'll get
	// [a1,a2,a3..., v, b1,b2...]
	arr[l+1], arr[len(arr)-1] = arr[len(arr)-1], arr[l+1]
	return l + 1
}
