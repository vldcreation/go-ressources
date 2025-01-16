package medium

func XorAllPairingNums(nums1 []int, nums2 []int) int {
	xor1, xor2, len1, len2 := 0, 0, len(nums1), len(nums2)
	if len1%2 != 0 {
		for _, num2 := range nums2 {
			xor2 ^= num2
		}
	}
	if len2%2 != 0 {
		for _, num1 := range nums1 {
			xor1 ^= num1
		}
	}
	return xor1 ^ xor2
}
