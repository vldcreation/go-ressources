package medium

func MinWindowSubstring(strArr [2]string) string {
	str := strArr[0]
	target := strArr[1]

	targetMap := make(map[rune]int)
	for _, v := range target {
		targetMap[v]++
	}

	strMap := make(map[rune]int)
	minStr := ""
	minLen := len(str) + 1
	left := 0

	// Extend the window
	for right := 0; right < len(str); right++ {
		strMap[rune(str[right])]++

		// Try to shrink the window while maintaining validity
		for left <= right && isValid(targetMap, strMap) {
			currLen := right - left + 1
			if currLen < minLen {
				minLen = currLen
				minStr = str[left : right+1]
			}

			// Try to shrink from left
			strMap[rune(str[left])]--
			left++
		}
	}

	return minStr
}

func isValid(targetMap, strMap map[rune]int) bool {
	for k, v := range targetMap {
		if strMap[k] < v {
			return false
		}
	}

	return true
}
