package maximum_subarray

import (
	"math"
)

func MaximumSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	_, _, sum := maximumSubarray(nums, 0, len(nums)-1)

	return sum
}

func maximumSubarray(nums []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, nums[low]
	}

	mid := (low + high) >> 1

	leftLow, leftHigh, leftSum := maximumSubarray(nums, low, mid)
	rightLow, rightHigh, rightSum := maximumSubarray(nums, mid+1, high)
	crossLow, crossHigh, crossSum := maxCrossingSubarray(nums, low, mid, high)

	switch {
	case leftSum >= rightSum && leftSum >= crossSum:
		return leftLow, leftHigh, leftSum
	case rightSum >= leftSum && rightSum >= crossSum:
		return rightLow, rightHigh, rightSum
	default:
		return crossLow, crossHigh, crossSum
	}
}

func maxCrossingSubarray(nums []int, low, mid, high int) (int, int, int) {
	maxLeft := mid
	leftSum, sum := math.MinInt, 0
	for i := mid; i >= low; i-- {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	maxRight := mid
	rightSum, sum := math.MinInt, 0
	for i := mid + 1; i <= high; i++ {
		sum += nums[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}
