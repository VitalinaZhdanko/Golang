package main

import (
	"fmt"
	"testing"
)

//TestDataItem struct for input data
type TestDataItem struct {
	start                 Point
	inputDate             Square
	side                  uint
	resEndPoint           Point
	resPerimeter, resArea uint
}

var dataItems = []TestDataItem{
	{inputDate: Square{Point{1, 1}, 5}, resEndPoint: Point{x: 6, y: 6}, resPerimeter: 20, resArea: 25},
	{inputDate: Square{Point{2, 2}, 6}, resEndPoint: Point{x: 8, y: 8}, resPerimeter: 24, resArea: 36},
}

func TestSquare(t *testing.T) {
	for _, item := range dataItems {
		res, err := item.inputDate.End()
		if err != nil {
			fmt.Println("Error...", err)
			return
		}
		if res != item.resEndPoint {
			t.Errorf("End() with args %v : FAILED, expected %v but got value '%v'", item.inputDate, item.resEndPoint, res)
		} else {
			t.Logf("End() with args %v : PASSED, expected %v and got value '%v'", item.inputDate, item.resEndPoint, res)
		}

	}
}

func TestPerimeter(t *testing.T) {
	for _, item := range dataItems {
		res, err := item.inputDate.Perimeter()
		if err != nil {
			fmt.Println("Error...", err)
			return
		}
		if res != item.resPerimeter {
			t.Errorf("Perimeter() with args %v : FAILED, expected %v but got value '%v'", item.side, item.resPerimeter, res)
		} else {
			t.Logf("Perimeter() with args %v : PASSED, expected %v and got value '%v'", item.side, item.resPerimeter, res)
		}

	}
}

func TestArea(t *testing.T) {
	for _, item := range dataItems {
		res, err := item.inputDate.Area()
		if err != nil {
			fmt.Println("Error...", err)
			return
		}
		if res != item.resArea {
			t.Errorf("Area() with args %v : FAILED, expected %v but got value '%v'", item.side, item.resArea, res)
		} else {
			t.Logf("Area() with args %v : PASSED, expected %v and got value '%v'", item.side, item.resArea, res)
		}

	}
}
