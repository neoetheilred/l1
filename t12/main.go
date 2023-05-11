package main

import "fmt"

/*
	A is a proper subset of B <=> all elements of A belong to B, A != B, A != {}

	In this task we'll assume that A is a proper subset of B <=> A != {}, A != B and exactly one element of B doesn't belong to A

	We'll use map[string]struct{} type for set of strings
*/

// Transforms slice of string to proper subset of strings
func ProperSubset(seq []string) (map[string]struct{}, bool) {
	if len(seq) <= 1 { // In this case there are no proper subsets
		return nil, false
	}
	res := make(map[string]struct{})
	// First, create a set from the sequense of strings
	for _, s := range seq {
		res[s] = struct{}{}
	}
	// Remove exactly one element
	for k := range res {
		delete(res, k)
		break
	}
	if len(res) == 0 { // If resulting set is emtpy, it is not a proper subset
		return nil, false
	}
	return res, true
}

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	seq2 := []string{"cat", "cat", "cat"}
	fmt.Println(ProperSubset(seq))
	fmt.Println(ProperSubset(seq2))
}
