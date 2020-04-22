package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

func print(figure Figure) {
	fmt.Println("Area: ", figure.area(), " Perimeter: ", figure.perimeter())
}

type Square struct{ side float64 }

func (s Square) area() float64 {
	return s.side * s.side
}
func (s Square) perimeter() float64 {
	return 4 * s.side
}

type Circle struct{ radius float64 }

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	s := Square{5}
	c := Circle{4}

	if s.side > 0 {
		fmt.Println("Square")
		print(s)
	}
	if c.radius > 0 {
		fmt.Println("Circle")
		print(c)
	}
}
