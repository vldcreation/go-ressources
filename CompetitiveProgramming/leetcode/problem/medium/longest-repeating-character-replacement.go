package medium

func CharacterReplacement(s string, k int) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	counts := [26]int{} // Frequency map for 'A' through 'Z'
	left := 0
	maxLength := 0
	maxFreq := 0 // Max frequency of *any single character* in the current window

	for right := 0; right < n; right++ {
		// Expand window by including s[right]
		currentCharIndex := s[right] - 'A'
		counts[currentCharIndex]++
		maxFreq = max(maxFreq, counts[currentCharIndex])

		// Current window size
		windowSize := right - left + 1

		// Check if the current window needs too many replacements
		// replacementsNeeded = windowSize - maxFreq
		if windowSize-maxFreq > k {
			// Shrink window from the left
			leavingCharIndex := s[left] - 'A'
			counts[leavingCharIndex]--
			left++
			// Note: We don't need to decrement maxFreq here.
			// The logic relies on maxFreq representing the max frequency seen
			// in *any* valid window considered so far. If we shrink,
			// the window size also shrinks, maintaining the balance for finding
			// the *overall* max length.
		}

		// Update maxLength. After the shrinking loop (if it ran),
		// the window [left, right] is the largest valid window ending at 'right'.
		// Or, if the shrinking loop didn't run, it was already valid.
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}
