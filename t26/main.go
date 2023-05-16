package main

/*
	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.
*/

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s := scanner.Text()
		fmt.Printf("All unique: %v\n", checkAllUnique(s))
	} else {
		fmt.Println("No input provided")
	}
}

func checkAllUnique(s string) bool {
	set := make(map[rune]struct{}) // Use map as set
	for _, c := range s {
		lower := unicode.ToLower(c)  // Cast unicode symbol to lowercase
		if _, ok := set[lower]; ok { // If such symbol already appeared => at least one symbol is repeated
			return false
		}
		set[lower] = struct{}{}
	}
	return true
}
