package medium

func FindClosestElements(arr []int, k int, x int) []int {
	ans := []int{}
	left, right := 0, len(arr)-1
	for right-left+1 > k {
		if abs(arr[left]-x) > abs(arr[right]-x) {
			left++
		} else {
			right--
		}
	}
	for i := left; i <= right; i++ {
		ans = append(ans, arr[i])
	}
	return ans
}
