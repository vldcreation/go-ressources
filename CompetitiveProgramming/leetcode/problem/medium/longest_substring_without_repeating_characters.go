package medium

func LengthOfLongestSubstring(s string) int {
	ans := 0
	bank := map[rune]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := bank[rune(s[i])]; ok {
			if i == len(s)-1 {
				return max(ans, len(bank))
			} else {
				ans = max(ans, len(bank))
				i = bank[rune(s[i])] + 1
				bank = make(map[rune]int)
				bank[rune(s[i])] = i
			}
			continue
		}

		bank[rune(s[i])] = i
	}

	return max(ans, len(bank))
}

func LengthOfLongestSubstring__GPT(s string) int {
	charIndex := make(map[rune]int) // Map to store the last index of each character
	maxLength := 0                  // Variable to store the maximum length found
	start := 0                      // Left boundary of the sliding window

	for i, char := range s {
		// If the character is already in the map and its index is within the current window
		if lastIndex, ok := charIndex[char]; ok && lastIndex >= start {
			start = lastIndex + 1 // Move the start to one position right of the last occurrence
		}
		charIndex[char] = i                   // Update or add the character's index
		maxLength = max(maxLength, i-start+1) // Calculate max length
	}

	return maxLength
}
