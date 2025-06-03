package maximum_subarray

import (
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "all positive numbers",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 15, // 1+2+3+4+5
		},
		{
			name:     "mixed positive and negative numbers",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6, // [4,-1,2,1]
		},
		{
			name:     "all negative numbers",
			nums:     []int{-5, -3, -8, -1, -2},
			expected: -1, // [-1] is the subarray with maximum sum
		},
		{
			name:     "single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "with zero",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 0},
			expected: 6, // [4,-1,2,1]
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaximumSubarray(tt.nums)
			if result != tt.expected {
				t.Errorf("MaximumSubarray(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
