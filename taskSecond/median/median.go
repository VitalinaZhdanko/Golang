package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {

	var number int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scan(&number)

	if number == 0 {
		fmt.Println("Enter correct data")
		return
	}
	slice := make([]int, number)
	for i := 0; i < number; i++ {
		slice[i] = rand.Intn(100)
	}
	fmt.Println(slice)
	fmt.Println(median(slice))

}

func median(slice []int) float64 {
	sort.Ints(slice)
	if len(slice)%2 == 0 {
		fmt.Println("Sorted even array: ", slice)
		return float64((slice[len(slice)/2] + slice[len(slice)/2-1])) / 2
	}
	fmt.Println("Sorted odd array: ", slice)
	return float64(slice[len(slice)/2])

}
