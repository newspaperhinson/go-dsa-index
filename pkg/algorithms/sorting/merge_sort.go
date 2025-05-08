package sorting

func MergeSort[S ~[]E, E any](arr S, lessThan func(a, b E) bool) {
	mergeSortInPlace(arr, 0, len(arr)-1, lessThan)
}

func mergeSortInPlace[S ~[]E, E any](arr S, left, right int, lessThan func(a, b E) bool) {
	if left >= right {
		return
	}

	mid := (left + right) / 2

	mergeSortInPlace(arr, left, mid, lessThan)
	mergeSortInPlace(arr, mid+1, right, lessThan)

	merge(arr, left, mid, right, lessThan)
}

func merge[S ~[]E, E any](arr S, left, mid, right int, lessThan func(a, b E) bool) {
	leftLength := mid - left + 1
	rightLength := right - mid

	leftArray := make(S, leftLength)
	copy(leftArray, arr[left:mid+1])
	rightArray := make(S, rightLength)
	copy(rightArray, arr[mid+1:right+1])

	cur := left
	i, j := 0, 0
	for i < leftLength && j < rightLength {
		if lessThan(leftArray[i], rightArray[j]) {
			arr[cur] = leftArray[i]
			i++
		} else {
			arr[cur] = rightArray[j]
			j++
		}
		cur++
	}

	for i < leftLength {
		arr[cur] = leftArray[i]
		i++
		cur++
	}

	for j < rightLength {
		arr[cur] = rightArray[j]
		j++
		cur++
	}
}
