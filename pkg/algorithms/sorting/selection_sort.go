package sorting

func SelectionSort[S ~[]E, E any](arr S, lessThan func(a, b E) bool) {
	for i := range len(arr) {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if lessThan(arr[j], arr[minIndex]) {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
