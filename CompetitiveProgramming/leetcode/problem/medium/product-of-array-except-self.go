package medium

import "log"

func ProductExceptSelf(nums []int) []int {
	left, right := make([]int, len(nums)), make([]int, len(nums))
	left[0] = 1
	right[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	out := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		log.Println(left[i], right[i])
		out[i] = left[i] * right[i]
	}

	return out
}

func ProductExceptSelfBest(nums []int) []int {
	out := make([]int, len(nums))
	out[0] = 1
	for i := 1; i < len(nums); i++ {
		out[i] = nums[i-1] * out[i-1]
	}

	R := 1
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] *= R
		R *= nums[i]
	}

	return out
}
