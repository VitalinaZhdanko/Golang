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
	} else {
		slice := make([]int, number)
		for i := 0; i < number; i++ {
			slice[i] = rand.Intn(100)
		}
		fmt.Println(slice)
		fmt.Println(median(slice))
	}
}

func median(slice []int) float64 {
	sort.Ints(slice)
	var num float64
	if len(slice)%2 == 0 {
		fmt.Println("Sorted even array: ", slice)
		num = float64((slice[len(slice)/2] + slice[len(slice)/2-1])) / 2
	} else {
		fmt.Println("Sorted odd array: ", slice)
		num = float64(slice[len(slice)/2])
	}
	return num
}
