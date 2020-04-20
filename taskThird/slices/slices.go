package main

import (
	"fmt"
)

func main() {
	sliceString1 := []string{"one", "two", "three"}
	sliceString2 := []string{"one", "two"}
	sliceString3 := []string{}
	fmt.Println("SliceString1: ", sliceString1)
	fmt.Println("The longest word from the sliceString1: ", max(sliceString1))
	fmt.Println("SliceString2: ", sliceString2)
	fmt.Println("The longest word from the sliceString2: ", max(sliceString2))
	fmt.Println("SliceString3: ", sliceString3)
	fmt.Println("The longest word from the sliceString3: ", max(sliceString3))

	sliceInt1 := []int64{1, 2, 5, 15}
	fmt.Println("SliceInt: ", sliceInt1)
	fmt.Println("Reverse sliceInt: ", reverse(sliceInt1))
	fmt.Println("SliceInt: ", sliceInt1)
	sliceInt2 := []int64{}
	fmt.Println("SliceInt: ", sliceInt2)
	fmt.Println("Reverse sliceInt: ", reverse(sliceInt2))
}

func max(slice []string) string {
	var maxWord string
	if len(slice) == 0 {
		maxWord = "Slice is empty!"
	} else {
		maxWord = slice[0]
		for _, value := range slice {
			if len(maxWord) < len(value) {
				maxWord = value
			}
		}
	}
	return maxWord
}

func reverse(slice []int64) []int64 {
	var reverseSlice []int64
	if len(slice) == 0 {
		fmt.Println("Slice is empty")
		reverseSlice = nil
	} else {
		for i := len(slice) - 1; i >= 0; i-- {
			reverseSlice = append(reverseSlice, slice[i])
		}
	}
	return reverseSlice
}
