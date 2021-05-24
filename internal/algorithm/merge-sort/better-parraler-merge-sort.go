package mergesort

import "sync"

func BetterParallelMergeSort(array []int) []int {
	ansChan := make(chan []int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(c chan []int) {
		BetterHParallelSimpleMergeSort(array, ansChan, 0)
		defer wg.Done()
	}(ansChan)
	ans := <-ansChan
	wg.Wait()
	return ans
}
func BetterHParallelSimpleMergeSort(array []int, ansChan chan []int, depth int) {
	defer close(ansChan)

	nubmerOfArray := len(array)
	if nubmerOfArray <= 1 {
		ansChan <- array
		return
	}

	if depth < 3 {
		wg := sync.WaitGroup{}
		wg.Add(2)
		sortedArrayLeftChan := make(chan []int)
		go func(sortedArrayLeftChan chan []int) {
			BetterHParallelSimpleMergeSort(
				array[nubmerOfArray/2:],
				sortedArrayLeftChan,
				depth+1,
			)
			defer wg.Done()
		}(sortedArrayLeftChan)

		sortedArrayRightChan := make(chan []int)
		go func(sortedArrayRightChan chan []int) {
			BetterHParallelSimpleMergeSort(
				array[:nubmerOfArray/2],
				sortedArrayRightChan,
				depth+1,
			)
			defer wg.Done()
		}(sortedArrayRightChan)
		ansChan <- Merge(<-sortedArrayLeftChan, <-sortedArrayRightChan)
	} else {
		ansChan <- Merge(
			SimpleMergeSort(array[nubmerOfArray/2:]),
			SimpleMergeSort(array[:nubmerOfArray/2]),
		)
	}

}
