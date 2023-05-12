package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	fmt.Printf("Got:\t\t%s\nReversed:\t%s", s, reverseString(s))
}

func reverseString(s string) string {
	return string(reverseSlice([]rune(s))) // Transform string to []rune in order to correctly proccess unicode
}

func reverseSlice[T any](a []T) []T {
	res := make([]T, len(a)) // Predefine length of resulting slice in order to avoid multiple memory allocations
	for i := 0; i < len(a); i++ {
		res[i] = a[len(a)-i-1]
	}
	return res
}
