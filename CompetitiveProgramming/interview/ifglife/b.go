package ifglife

/* FindShortestSubstring
   Given a string. The tas is to find the length of the shortest substring,
   which upon deletion, makes the resultant string to be consisting
   of distinct characters only.

   A substring is a contiguous sequence of characters within a string.
   When a substring is deleted, one needs to merge the rest of
   the character blocks of the string(s).
   If no substring needs to be deleted, the answer is 0.

   Consider the given string to be , delete the substring "bb" in the range [3, 4]
   to get the remaining string "abck" which consists of distinct characters only.
   This is the minimum possible length of the string. Hence, the answer is 2.
*/

func FindShortestSubstring(s string) int {
	n := len(s)

	// If string is empty or has length 1, no need to remove anything
	if n <= 1 {
		return 0
	}

	// Create frequency array instead of a map for better performance
	freq := make([]int, 256)
	duplicates := 0

	// Count initial frequencies
	for _, c := range s {
		freq[c]++
		if freq[c] == 2 {
			duplicates++
		}
	}

	// If no duplicates exist
	if duplicates == 0 {
		return 0
	}

	minLength := n

	// Use two pointers approach
	for start := 0; start < n; start++ {
		tempFreq := make([]int, 256)
		copy(tempFreq, freq)
		tempDuplicates := duplicates

		for end := start; end < n && tempDuplicates > 0; end++ {
			c := s[end]
			tempFreq[c]--
			if tempFreq[c] == 1 {
				tempDuplicates--
			}

			if tempDuplicates == 0 {
				minLength = min(minLength, end-start+1)
				break
			}
		}

		// Early termination if we found a substring of length 1
		if minLength == 1 {
			break
		}
	}

	return minLength
}
