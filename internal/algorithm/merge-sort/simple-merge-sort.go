package mergesort

func SimpleMergeSort(array []int) []int {

	nubmerOfArray := len(array)
	if nubmerOfArray <= 1 {
		return array
	}

	return Merge(
		SimpleMergeSort(array[nubmerOfArray/2:]),
		SimpleMergeSort(array[:nubmerOfArray/2]),
	)
}
func Merge(arrayLeft, arrayRight []int) []int {

	numberOfLeftArray := len(arrayLeft)
	numberOfRightArray := len(arrayRight)
	numerberOfMergedArray := numberOfLeftArray + numberOfRightArray

	arrayMerged := make([]int, 0, numerberOfMergedArray)

	pLeft, pRright := 0, 0

	for pLeft < numberOfLeftArray && pRright < numberOfRightArray {
		if arrayLeft[pLeft] < arrayRight[pRright] {
			arrayMerged = append(arrayMerged, arrayLeft[pLeft])
			pLeft++
		} else {
			arrayMerged = append(arrayMerged, arrayRight[pRright])
			pRright++
		}
	}
	if pLeft < numberOfLeftArray {
		arrayMerged = append(arrayMerged, arrayLeft...)
	}
	if pRright < numberOfRightArray {
		arrayMerged = append(arrayMerged, arrayRight...)
	}
	return arrayMerged
}
