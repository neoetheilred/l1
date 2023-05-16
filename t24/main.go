package main

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

import (
	"fmt"
	"math"
)

func main() {
	a, b := NewPoint(0, 0), NewPoint(1, 1)
	fmt.Println(a.Dist(b))
}

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (a Point) Dist(b Point) float64 {
	return math.Sqrt(
		(a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y),
	)
}
