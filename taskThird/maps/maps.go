package main

import (
	"fmt"
	"sort"
)

func main() {
	mapOne := map[int]string{2: "a", 0: "b", 1: "c"}
	mapTwo := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	fmt.Println("mapOne: ", mapOne)
	printSorted(mapOne)
	fmt.Println("mapTwo: ", mapTwo)
	printSorted(mapTwo)
}

func printSorted(myMap map[int]string) {
	keys := make([]int, 0, len(myMap))

	for key := range myMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	fmt.Print("Sorted map by key: ")
	for _, k := range keys {
		fmt.Print(myMap[k] + " ")
	}
	fmt.Println()
}
