package main

import (
	"fmt"

	mergesort "github.com/ahmadrezam97/hand-to-hand/internal/algorithm/merge-sort"
)

func main() {
	array := []int{11, 1, 13, 14, 15, 15, 1, 2, 3, 4, 5}
	sortedArray := mergesort.BetterMergeSort(array)
	fmt.Println(sortedArray)
}
