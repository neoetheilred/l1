package main

/*
	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/

import "fmt"

func main() {
	s := "a"
	i := 5
	b := false
	var ch chan struct{}
	PrintType(s)
	PrintType(i)
	PrintType(b)
	PrintType(ch)
	PrintType(struct{ a int }{5})
}

// Uses type switch to check the types in runtime
func PrintType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("V is integer")
	case string:
		fmt.Println("V is string")
	case bool:
		fmt.Println("V is boolean")
	case chan struct{}:
		fmt.Println("V is channel of struct{}")
	default:
		fmt.Printf("V is of type %T", v)
	}
}
