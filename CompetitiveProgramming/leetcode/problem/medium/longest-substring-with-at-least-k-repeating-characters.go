package medium

import "strings"

func LongestSubstring(s string, k int) int {
	if len(s) == 0 || len(s) < k {
		return 0
	}

	// Count frequency of each character
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// Find characters that appear less than k times
	var splitChar byte
	for char, count := range freq {
		if count < k {
			splitChar = char
			break
		}
	}

	// If no character appears less than k times, entire string is valid
	if freq[splitChar] == 0 {
		return len(s)
	}

	// Split string at characters that appear less than k times and recurse
	substrings := strings.Split(s, string(splitChar))
	maxLen := 0
	for _, substr := range substrings {
		maxLen = max(maxLen, LongestSubstring(substr, k))
	}

	return maxLen
}
