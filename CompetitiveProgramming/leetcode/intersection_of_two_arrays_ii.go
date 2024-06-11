package leetcode

func check(nums1, nums2 []int) ([]int, []int) {
	if len(nums1) <= len(nums2) {
		return nums1, nums2
	}

	return nums2, nums1
}

func intersect(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0)
	fIdx := 1
	nums1, nums2 = check(nums1, nums2)
	if len(nums1) <= 1 {
		return nums1
	}

	for k1, n1 := range nums1 {
		for k2, n2 := range nums2 {
			if len(ans) < len(nums1) {
				if k1 == 0 && k2 == 0 && n1 == n2 {
					ans = append(ans, n1)
					break
				} else {
					if n1 == n2 && fIdx != k2 {
						ans = append(ans, n1)
						fIdx = k2
						break
					}
				}
			}
		}
	}

	return ans
}
