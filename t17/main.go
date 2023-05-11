package main

import "fmt"

func main() {
	t := []int{1, 3, 5, 7, 8, 9, 10, 10, 10}
	fmt.Println(binSearch(t, 10)) // 6
	fmt.Println(binSearch(t, 6))  // 3
}

// Returns index of left-most element `e`, e >= n
func binSearch(arr []int, n int) (index int) {
	l, r := -1, len(arr)
	for l < r-1 {
		m := (l + r) / 2
		if arr[m] < n {
			l = m
		} else {
			r = m
		}
	}
	return r
}
