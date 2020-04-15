package main

import (
	"fmt"
	"math"
)

func add(x float64, y float64) float64 {
	return math.Pi + x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(x int) (a, b int) {
	a = x + 1
	b = x - 1
	return
}

var an, bn, cn bool

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(add(1, 2))
	fmt.Println(math.Pi)
	fmt.Println(swap("hello", "world"))
	a, b := swap("hi", "vita")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(split(10))
	var dn, en int
	fmt.Println(an, bn, cn, dn, en)
}
