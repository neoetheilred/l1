package main

import (
	"fmt"
)

// Type for Set of elements (we could directly use map[T]struct, but defining new type is more convenient)
type Set[T comparable] struct {
	m map[T]struct{}
}

// Constructor
func NewSet[T comparable](vals ...T) *Set[T] {
	var s *Set[T] = &Set[T]{m: make(map[T]struct{})}
	for _, v := range vals {
		s.m[v] = struct{}{}
	}
	return s
}

// Returns new set, which is intersection of sets `a` and `b`
func (a *Set[T]) Intersect(b *Set[T]) *Set[T] {
	res := NewSet[T]()
	for e := range a.m {
		_, ok := b.m[e]
		if ok {
			res.m[e] = struct{}{}
		}
	}
	return res
}

// Retuns element count in set
func (a *Set[T]) Size() int {
	return len(a.m)
}

// Returns elements of set as a slice (not sorted)
func (a *Set[T]) Elements() []T {
	elements := make([]T, a.Size())
	i := 0
	for k := range a.m {
		elements[i] = k
		i++
	}
	return elements
}

// Implementation of Stringer interface
func (a *Set[T]) String() string {
	buf := []byte("{")
	for k := range a.m {
		buf = append(buf, []byte(fmt.Sprintf("%v, ", k))...) // Efficient concatenation
	}
	if a.Size() > 0 {
		buf = buf[:len(buf)-2]
	}
	buf = append(buf, '}')
	return string(buf)
}

func main() {
	a := NewSet(1, 2, 3, 4, 5)
	b := NewSet(2, 4, 6, 8, 10, 5)
	fmt.Println(a.Intersect(b))
}
