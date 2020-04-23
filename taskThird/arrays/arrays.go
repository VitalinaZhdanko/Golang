package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("Array: ", array)
	fmt.Println("Average value: ", average(array))
}

func average(array [6]int) float64 {
	var sum int
	for _, value := range array {
		sum += value
	}
	return float64(sum) / float64(len(array))
}
