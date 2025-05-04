package sorting

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name     string
	input    []int
	expected []int
}

var testCases = []TestCase{
	{
		name:     "Already sorted array",
		input:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Reverse sorted array",
		input:    []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "Array with duplicate elements",
		input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5},
		expected: []int{1, 1, 2, 3, 4, 5, 5, 6, 9},
	},
	{
		name:     "Empty array",
		input:    []int{},
		expected: []int{},
	},
	{
		name:     "Single element array",
		input:    []int{42},
		expected: []int{42},
	},
	{
		name:     "Negative numbers",
		input:    []int{-5, 3, -2, 0, 7, -1},
		expected: []int{-5, -2, -1, 0, 3, 7},
	},
}

func TestSort(t *testing.T) {
	for _, tc := range testCases {
		input := make([]int, len(tc.input))
		copy(input, tc.input)

		InsertionSort(input)

		if !reflect.DeepEqual(input, tc.expected) {
			t.Errorf("Test case %s failed: got %v, want %v", tc.name, input, tc.expected)
		}
	}
}
