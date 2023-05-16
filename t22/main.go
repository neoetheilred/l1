package main

/*
	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

import (
	"fmt"
	"math/big"
)

func main() {
	bigOne := big.NewInt(1)
	a, b := big.NewInt(0).Lsh(bigOne, 500), big.NewInt(0).Lsh(bigOne, 300) // a = 1 << 500, b = 1 << 300
	fmt.Printf("a = %v\nb = %v\n\n", a, b)
	fmt.Printf("a + b = %v\na * b = %v\na - b = %v\na / b = %v\n",
		big.NewInt(0).Add(a, b),
		big.NewInt(0).Mul(a, b),
		big.NewInt(0).Sub(a, b),
		big.NewInt(0).Div(a, b),
	)

}
