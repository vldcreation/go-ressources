package easy

// 1, 1, 0, 1, 1, 1
// return 3
// 1, 0, 1, 1, 0, 1
// return 2
func FindMaxConsecutiveOnes(nums []int) int {
	ans, count := 0, 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			ans = max(ans, count)
			count = 0
		}
	}
	return max(ans, count)
}

func FindMaxConsecutiveOnesV2(nums []int) int {
	left, right, maxCount := 0, 0, 0

	for right < len(nums) {
		if nums[right] == 1 {
			right++
		} else {
			maxCount = max(maxCount, right-left)
			right++
			left = right
		}
	}
	maxCount = max(maxCount, right-left)
	return maxCount
}
