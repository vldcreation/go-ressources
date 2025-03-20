package medium

func isPalindrome(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func best(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}

	return s2
}

func LongestPalindrome(s string) string {
	var ans string

	if len(s) < 3 {
		if isPalindrome(s) {
			return s
		}
		return string(s[0])
	}

	curStr := ""
	startIdx := 0
	for i := 0; i < len(s); i++ {
		curStr = curStr + string(s[i])
		if isPalindrome(curStr) {
			ans = best(ans, curStr)
		} else {
			// reset
			if i == len(s)-1 {
				if startIdx <= len(s)-1 {
					startIdx = startIdx + 1
					i = startIdx
				}
				curStr = string(s[startIdx])
			}
		}
	}

	if ans == "" {
		return string(s[0])
	}

	return ans
}

func LongestPalindromeBest(s string) string {
	if len(s) < 2 {
		return s
	}

	start, maxLen := 0, 1

	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			curLen := right - left + 1
			if curLen > maxLen {
				start = left
				maxLen = curLen
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expandAroundCenter(i, i)   // Odd length palindromes
		expandAroundCenter(i, i+1) // Even length palindromes
	}

	return s[start : start+maxLen]
}
