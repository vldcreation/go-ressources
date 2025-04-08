package medium

func NumberOfArithmeticSlices(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	ans := 0
	dp := make([]int, len(nums))
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
			ans += dp[i]
		}
	}
	return ans
}
