package mergesort

import "sync"



func ParallelMergeSort(array []int )[]int{
	ansChan := make(chan []int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func (c chan []int)  {
		HParallelSimpleMergeSort(array, ansChan)
		defer wg.Done()
	}(ansChan)
	ans := <-ansChan
	wg.Wait()
	return ans
}
func HParallelSimpleMergeSort(array []int, ansChan chan []int) {
	defer close(ansChan)

	nubmerOfArray := len(array)
	if nubmerOfArray <= 1 {
		ansChan <- array
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	sortedArrayLeftChan := make(chan []int)
	go func(sortedArrayLeftChan chan []int) {
		HParallelSimpleMergeSort(
			array[nubmerOfArray/2:],
			sortedArrayLeftChan,
		)
		defer wg.Done()
	}(sortedArrayLeftChan)

	sortedArrayRightChan := make(chan []int)
	go func(sortedArrayRightChan chan []int) {
		HParallelSimpleMergeSort(
			array[:nubmerOfArray/2],
			sortedArrayRightChan,
		)
		defer wg.Done()
	}(sortedArrayRightChan)

	ansChan <- Merge(<-sortedArrayLeftChan, <-sortedArrayRightChan)
}
