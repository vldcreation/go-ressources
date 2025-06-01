package easy

func CountGoodSubstrings(s string) int {
	ans := 0
	for i := 2; i < len(s); i++ {
		if s[i] != s[i-1] && s[i-1] != s[i-2] && s[i] != s[i-2] {
			ans++
		}
	}

	return ans
}
