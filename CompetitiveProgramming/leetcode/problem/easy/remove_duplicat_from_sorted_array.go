package easy

func RemoveDuplicatesFromSortedArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize two pointers
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			if j > i {
				nums[i] = nums[j]
			}
		}
	}

	return i + 1
}
