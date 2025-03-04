package medium

func MinSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1
	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			minLen = min(minLen, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}

	return minLen
}
