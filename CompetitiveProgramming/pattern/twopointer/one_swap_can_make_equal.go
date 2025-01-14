package twopointer

func CheckIfOneSwapCanMakeEqual(s1, s2 string) bool {
	// If the strings are equal, return true
	if s1 == s2 {
		return true
	}

	// Find the first and last index where the strings differ
	firstDiff, lastDiff := -1, -1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if firstDiff == -1 {
				firstDiff = i
			}
			lastDiff = i
		}
	}

	// If there are no differences, return false
	if firstDiff == -1 {
		return false
	}

	// Swap the characters and check if the strings are equal
	return s1[firstDiff] == s2[lastDiff] && s1[lastDiff] == s2[firstDiff]
}
