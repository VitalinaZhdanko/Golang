package main

import (
	"errors"
	"fmt"
)

//Point struct
type Point struct {
	x, y int
}

//Square struct with item Point
type Square struct {
	start Point
	a     uint
}

func main() {
	s := Square{Point{-1, -1}, 5}
	fmt.Println("Side length: ", s.a)
	fmt.Println("Starting points: ", s.start)

	resEndPoint, err := s.End()
	if err != nil {
		fmt.Println("Error...", err)
		return
	}
	fmt.Println("End point: ", resEndPoint)

	resPerimeter, err := s.Perimeter()
	if err != nil {
		fmt.Println("Error...", err)
		return
	}
	fmt.Println("Perimeter: ", resPerimeter)

	resArea, err := s.Area()
	if err != nil {
		fmt.Println("Error...", err)
		return
	}
	fmt.Println("Area: ", resArea)

}

//End method for Square
func (s Square) End() (Point, error) {
	if s.a >= 0 {
		return Point{s.start.x + int(s.a), s.start.y + int(s.a)}, nil
	}
	return Point{}, errors.New("No correct data")

}

//Perimeter for Square
func (s Square) Perimeter() (uint, error) {
	if s.a > 0 {
		return 4 * s.a, nil
	}
	return 0, errors.New("No correct data")
}

//Area for Square
func (s Square) Area() (uint, error) {
	if s.a > 0 {
		return s.a * s.a, nil
	}
	return 0, errors.New("No correct data")
}
