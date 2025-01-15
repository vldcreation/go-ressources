package easy

func RemoveDuplicatesFromSortedArray(nums []int) int {
	mp := make(map[int]int)

	for _, n := range nums {
		if _, ok := mp[n]; !ok {
			mp[n]++
		}
	}

	nums = nums[:0]
	for _, n := range mp {
		nums = append(nums, n)
	}

	return len(mp)
}
