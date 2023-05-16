package main

/*
	Удалить i-ый элемент из слайса
*/

import (
	"fmt"
	"math/rand"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7} // Fake testing data
	i := rand.Intn(len(arr))
	fmt.Printf("slice: %v\nindex to remove at: %d\n", arr, i)
	removeAt(&arr, i)
	fmt.Printf("slice after removal: %v\n\n", arr)
	i = rand.Intn(len(arr))
	fmt.Printf("slice: %v\nindex to remove at: %d\n", arr, i)
	removeAtOrderSafe(&arr, i)
	fmt.Printf("slice after order-safe removal: %v", arr)
}

func removeAtOrderSafe(arr *[]int, i int) {
	*arr = append((*arr)[:i], (*arr)[i+1:]...)
}

// Removes element at i-th position at constant time complexity
// BREAKS element order
func removeAt(arr *[]int, i int) {
	if i < 0 || len(*arr) <= i { // Check if given index is correct
		return // Cannot remove at invalid index, do nothing
		// We could panic or return error instead
	}

	(*arr)[i] = (*arr)[len(*arr)-1]
	(*arr) = (*arr)[:len(*arr)-1]
}
