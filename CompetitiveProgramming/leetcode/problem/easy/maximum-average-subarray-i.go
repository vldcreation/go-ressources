package easy

func FindMaxAverage(nums []int, k int) float64 {
	left := 0
	right := k
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	mx := sum
	for right < len(nums) {
		sum -= nums[left]
		left++
		sum += nums[right]
		right++
		mx = max(mx, sum)
	}
	return float64(mx) / float64(k)
}
