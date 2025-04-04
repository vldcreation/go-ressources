package medium

func LongestSubarray(nums []int) int {
	left, zeros, maxLen := 0, 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros++
		}

		// Shrink window if we have more than one zero
		for zeros > 1 {
			if nums[left] == 0 {
				zeros--
			}
			left++
		}

		// Update maxLen if current window is larger
		currLen := right - left + 1
		if currLen > maxLen {
			maxLen = currLen
		}
	}

	// If we haven't found any zeros, we must delete one element
	if zeros == 0 && len(nums) > 0 {
		return len(nums) - 1
	}

	// Subtract 1 to account for the deleted zero
	return maxLen - 1
}
