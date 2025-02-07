package medium

func LetterCombinations(digits string) []string {
	var mapAlphabet = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return mapAlphabet[digits]
	}

	result := []string{}
	for _, c := range mapAlphabet[digits[:1]] {
		for _, s := range LetterCombinations(digits[1:]) {
			result = append(result, c+s)
		}
	}
	return result
}
