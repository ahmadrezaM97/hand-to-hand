package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"

	mergesort "github.com/ahmadrezam97/hand-to-hand/internal/algorithm/merge-sort"
)

func RunWithTimeMeasurement(mergeSortImp func([]int) []int, array []int) int64 {

	startTime := time.Now()
	mergeSortImp(array)
	return time.Now().Sub(startTime).Milliseconds()

}

func main() {

	array := []int{}
	for i := 0; i < 1000*100; i++ {
		array = append(array, rand.Intn(1000*1000))
	}
	mergeSortImpList := []func([]int) []int{
		mergesort.SimpleMergeSort,
		mergesort.ParallelMergeSort,
		mergesort.BetterMergeSort,
		mergesort.BetterParallelMergeSort,
	}
	for _, mergeSortImp := range mergeSortImpList {
		fmt.Println(
			runtime.FuncForPC(reflect.ValueOf(mergeSortImp).Pointer()).Name(),
			RunWithTimeMeasurement(mergeSortImp, array),
		)
	}
}
