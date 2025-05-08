package sorting

func InsertionSort[S ~[]E, E any](arr S, lessThan func(a, b E) bool) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]

		j := i - 1

		for j >= 0 && lessThan(key, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}
