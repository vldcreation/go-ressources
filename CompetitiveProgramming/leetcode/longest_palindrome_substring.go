package leetcode

func isPalindrome(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func best(s1, s2 string) string {
	if s1 > s2 {
		return s1
	}

	return s2
}

func longestPalindrome(s string) string {
	var ans string

	if len(s) < 3 {
		if isPalindrome(s) {
			return s
		}
		return string(s[0])
	}

	curStr := string(s[0])
	startIdx := 1
	for i := 1; i <= len(s); i++ {
		if !isPalindrome(curStr) && (len(ans) >= len(s)-i) {
			return ans
		}

		curStr = curStr + string(s[i])
		if isPalindrome(curStr) {
			ans = best(ans, curStr)

			// reset
			curStr = string(s[startIdx])
			if startIdx <= len(s)-1 {
				startIdx = startIdx + 1
			}
			i = startIdx
		}

		if i == len(s)-1 {
			curStr = string(s[startIdx])
			i = startIdx
			if startIdx <= len(s)-1 {
				startIdx = startIdx + 1
			}
		}
	}

	if ans == "" {
		return string(s[0])
	}

	return ans
}
