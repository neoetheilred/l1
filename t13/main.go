package main

import "fmt"

func main() {
	a, b := 10, 4520

	swap(&a, &b)

	fmt.Println(a, b)

	goswap(&a, &b)
	fmt.Println(a, b)
}

// Uses bitwise xor operator
// Since x ^ x == 0, 0 ^ x == x and xor is commutative, associative operation consider following:
// given a, b
// a = a ^ b
// b = b ^ a -> b ^ a ^ b -> a
// a = a ^ b -> a ^ b ^ a -> b
func swap(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func goswap(a, b *int) {
	*a, *b = *b, *a
}
