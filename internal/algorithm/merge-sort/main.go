package mergesort

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

//[)
func MeregBetter(array []int, l, r int, tmpArray []int) {
	m := (l + r) / 2
	i := l
	j := m
	cnt := l
	for i < m && j < r {
		if array[i] < array[j] {
			tmpArray[cnt] = array[i]
			i++
			cnt++
		} else {
			tmpArray[cnt] = array[j]
			j++
			cnt++
		}
	}

	for i < m {
		tmpArray[cnt] = array[i]
		i++
		cnt++
	}
	for j < r {
		tmpArray[cnt] = array[j]
		j++
		cnt++
	}

	for ind := l; ind < r; ind++ {
		array[ind] = tmpArray[ind]
	}
}

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

func BetterMergeSort(array []int) []int {

	lenght := len(array)
	nArray := make([]int, 0, lenght)
	tmpArray := make([]int, lenght, lenght)
	nArray = append(nArray, array...)

	HBMergeSort(nArray, 0, lenght, tmpArray)

	return nArray
}

func HBMergeSort(array []int, l, r int, tmpArray []int) {
	if r-l <= 1 {
		return
	}
	mid := (l + r) / 2
	HBMergeSort(array, l, mid, tmpArray)
	HBMergeSort(array, mid, r, tmpArray)
	MeregBetter(array, l, r, tmpArray)

}
