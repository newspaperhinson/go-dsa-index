package sorting

func InsertionSort[T any](arr []T, compare func(a, b T) bool) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]

		j := i - 1

		for j >= 0 && compare(key, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}
