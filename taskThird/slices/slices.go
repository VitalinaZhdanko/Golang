package main

import (
	"errors"
	"fmt"
)

func main() {
	sliceString1 := []string{"one", "two", "three"}
	sliceString2 := []string{"one", "two"}
	sliceString3 := []string{}

	printStringSlice(sliceString1)
	printStringSlice(sliceString2)
	printStringSlice(sliceString3)

	sliceInt1 := []int64{1, 2, 5, 15}
	sliceInt2 := []int64{}

	printIntSlice(sliceInt1)
	printIntSlice(sliceInt2)
}

func printStringSlice(slice []string) {
	fmt.Println("SliceString: ", slice)
	max, err := max(slice)
	if err != nil {
		fmt.Println("Error...", err)
		return
	}
	fmt.Println("The longest word from the sliceString: ", max)
}

func printIntSlice(slice []int64) {
	fmt.Println("SliceInt: ", slice)
	reverse, err := reverse(slice)
	if err != nil {
		fmt.Println("Error...", err)
		return
	}
	fmt.Println("Reverse sliceInt: ", reverse)
	fmt.Println("SliceInt: ", slice)
}

func max(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("Slice is empty")
	}
	var maxWord string = slice[0]
	for _, value := range slice {
		if len(maxWord) < len(value) {
			maxWord = value
		}
	}
	return maxWord, nil
}

func reverse(slice []int64) ([]int64, error) {
	var reverseSlice []int64
	if len(slice) == 0 {
		return nil, errors.New("Slice is empty")
	}
	for i := len(slice) - 1; i >= 0; i-- {
		reverseSlice = append(reverseSlice, slice[i])
	}
	return reverseSlice, nil
}
