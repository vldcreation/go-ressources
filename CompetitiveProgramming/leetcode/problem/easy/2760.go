package easy

func LongestAlternatingSubarray(nums []int, threshold int) int {
	n := len(nums)
	ans := 0

	for l := 0; l < n; l++ {
		// Check conditions for the start of the subarray (nums[l]):
		// 1. nums[l] % 2 == 0
		// 2. nums[l] <= threshold
		if nums[l]%2 == 0 && nums[l] <= threshold {
			// If we are here, nums[l] itself forms a valid subarray of at least length 1.
			currentLength := 1
			if currentLength > ans {
				ans = currentLength
			}

			// Try to extend the subarray from l+1.
			// r is the potential next element in the subarray.
			for r := l + 1; r < n; r++ {
				// Check conditions for nums[r] and the relationship with nums[r-1]:
				// 1. nums[r] <= threshold
				// 2. nums[r-1] % 2 != nums[r] % 2 (alternating parity)
				if nums[r] <= threshold && nums[r-1]%2 != nums[r]%2 {
					currentLength++
					if currentLength > ans {
						ans = currentLength
					}
				} else {
					// If conditions are not met, the subarray starting at l cannot be extended further.
					break
				}
			}
		}
		// If nums[l] is not a valid start, we move to the next potential start l+1.
	}

	return ans
}
