package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigOne := big.NewInt(1)
	a, b := big.NewInt(0).Lsh(bigOne, 500), big.NewInt(0).Lsh(bigOne, 300)
	// a, b := big.NewInt(3), big.NewInt(5)
	fmt.Printf("a = %v\nb = %v\n\n", a, b)
	// sum := (&big.Int{}).Add(a, b)
	// fmt.Println(sum)
	fmt.Printf("a + b = %v\na * b = %v\na - b = %v\na / b = %v\n",
		big.NewInt(0).Add(a, b),
		big.NewInt(0).Mul(a, b),
		big.NewInt(0).Sub(a, b),
		big.NewInt(0).Div(a, b),
	)

}
