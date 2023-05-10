package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1]) // Get i from command-line args, i-th bit will be set to 1
	if err != nil || n < 0 || n > 63 {
		log.Fatal("Command-line args: <N:integer from 0 to 63>")
	}
	res := int64(1) << n
	fmt.Println(res)
}
