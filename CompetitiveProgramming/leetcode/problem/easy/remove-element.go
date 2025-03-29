package easy

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	// initialize two pointers
	i, j := 0, len(nums)-1
	// loop until i and j meet
	for i <= j {
		// if nums[i] is equal to val, swap it with nums[j]
		if nums[i] == val {
			swap(&nums[i], &nums[j])
			j--
		} else {
			i++
		}
	}
	return i
}

func swap(a, b *int) {
	*a, *b = *b, *a
}
