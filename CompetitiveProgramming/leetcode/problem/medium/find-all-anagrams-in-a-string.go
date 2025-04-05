package medium

func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	ans := []int{}
	pCount := make(map[byte]int)
	windowCount := make(map[byte]int)

	// Count characters in pattern
	for i := 0; i < len(p); i++ {
		pCount[p[i]]++
	}

	// Initialize window
	for i := 0; i < len(p); i++ {
		windowCount[s[i]]++
	}

	// Check first window
	if mapsEqual(pCount, windowCount) {
		ans = append(ans, 0)
	}

	// Slide window
	for i := len(p); i < len(s); i++ {
		// Remove leftmost character
		windowCount[s[i-len(p)]]--
		if windowCount[s[i-len(p)]] == 0 {
			delete(windowCount, s[i-len(p)])
		}

		// Add rightmost character
		windowCount[s[i]]++

		// Check if current window is an anagram
		if mapsEqual(pCount, windowCount) {
			ans = append(ans, i-len(p)+1)
		}
	}

	return ans
}

func mapsEqual(m1, m2 map[byte]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
