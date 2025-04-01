package medium

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	ans := math.MaxInt
	diff := math.MaxInt

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				sum := nums[i] + nums[j] + nums[k]
				currDiff := abs(sum - target)
				if currDiff < diff {
					diff = currDiff
					ans = sum
				}
			}
		}
	}

	return ans
}

func ThreeSumClosestOptimized(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt
	ans := 0
	n := len(nums)

	for i := 0; i < n-2; i++ {

		// init left and right pointers
		left, right := i+1, n-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff := abs(sum - target)

			if diff < minDiff {
				minDiff = diff
				ans = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return ans
}
