package easy

func FindLHS(nums []int) int {
	counts := make(map[int]int)
	maxLength := 0

	for _, num := range nums {
		counts[num]++
	}

	for num, count := range counts {
		if countPlusOne, exists := counts[num+1]; exists {
			maxLength = max(maxLength, count+countPlusOne)
		}
	}

	return maxLength
}
