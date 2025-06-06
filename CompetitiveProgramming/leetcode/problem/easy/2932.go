package easy

func MaximumStrongPairXor(nums []int) int {
	maxXor := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if max(nums[i], nums[j])-min(nums[i], nums[j]) <= min(nums[i], nums[j]) {
				xorValue := nums[i] ^ nums[j]
				if xorValue > maxXor {
					maxXor = xorValue
				}
			}
		}
	}

	return maxXor
}
