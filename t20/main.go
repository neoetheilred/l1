package main

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
		fmt.Println(reverseSlice(split(s, ' ')))
	} else {
		fmt.Println("No input provided")
	}
}

func split(s string, delim rune) []string {
	res := []string{}     // Resulting slice
	t := []rune{}         // Use it for effecient string construction
	for _, c := range s { // c is rune, works fine with unicode
		if c == delim { // When we face the delimeter, we may want to append constructed string to resulting slice
			if len(t) != 0 { // Append only if `t` is not empty (for example, if there are multiple whitespaces)
				res = append(res, string(t))
				t = t[:0] // do not recreate slice, avoiding multiple memory allocations
			}
		} else { // Current character is not a delimeter
			t = append(t, c) // Add current character to string which is being constructed
		}
	}
	if len(t) != 0 { // Add last element, if any
		res = append(res, string(t))
	}
	return res
}

func reverseSlice[T any](a []T) []T {
	res := make([]T, len(a)) // Predefine length of resulting slice in order to avoid multiple memory allocations
	for i := 0; i < len(a); i++ {
		res[i] = a[len(a)-i-1]
	}
	return res
}
