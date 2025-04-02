package easy

func ContainsNearbyDuplicate(nums []int, k int) bool {
	i, j := 0, 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			if j-i <= k {
				return true
			}
		}
		if j-i == k {
			i++
		}
		j++
	}

	return false
}
