package main

import (
	"fmt"
)

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println("Side length: ", s.a)
	fmt.Println("Starting points: ", s.start)
	fmt.Println("End point: ", s.End())
	fmt.Println("Perimeter: ", s.Perimeter())
	fmt.Println("Area: ", s.Area())
}

func (s Square) End() Point {
	return Point{s.start.x + int(s.a), s.start.y + int(s.a)}
}
func (s Square) Perimeter() uint {
	return 4 * s.a
}
func (s Square) Area() uint {
	return s.a * s.a
}
