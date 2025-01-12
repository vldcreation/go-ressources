package medium

func FindMaxConsecutiveOnesIII(nums []int, k int) int {
	left, right, maxOnesLength := 0, 0, 0

	// Iterate through the array with the right pointer
	for right < len(nums) {
		// If we encounter a 0, decrement k (the flip count)
		if nums[right] == 0 {
			k--
		}

		// If k is negative, it means we've flipped more 0s than allowed
		for k < 0 {
			// If the left element is 0, we increment k
			// since we are moving past the flipped zero
			if nums[left] == 0 {
				k++
			}
			// Move the left pointer to the right, effectively shrinking the window
			left++
		}

		// Update maxOnesLength if the current window is larger
		maxOnesLength = max(maxOnesLength, right-left+1)

		// Move the right pointer to the next element
		right++
	}

	// Return the maximum length of the subarray
	return maxOnesLength
}
