package easy

import (
	"math"
	"sort"
)

func MinimumDifference(nums []int, k int) int {
	// If k is 1, the difference between the highest and lowest score is 0,
	// as we pick only one score, and score - score = 0.
	if k == 1 {
		return 0
	}

	// Sort the array to easily find min and max in a window.
	sort.Ints(nums) // Sorts in ascending order

	minDifference := math.MaxInt32

	// Use a sliding window of size k.
	// The loop runs from the first possible start of a window (index 0)
	// to the last possible start of a window (index len(nums)-k).
	for i := 0; i <= len(nums)-k; i++ {
		// The current window of k scores is from nums[i] to nums[i+k-1].
		// Since the array is sorted:
		// - nums[i] is the lowest score in this window.
		// - nums[i+k-1] is the highest score in this window.
		difference := nums[i+k-1] - nums[i]

		// Update the minimum difference found so far.
		if difference < minDifference {
			minDifference = difference
		}
	}

	return minDifference
}
