package easy

import "unicode"

func LongestNiceSubstring(s string) string {
	sset := make(map[rune]bool)
	for _, r := range s {
		sset[r] = true
	}

	off := 0
	is_full := true
	for i := range s {
		if !((unicode.IsLower(rune(s[i])) && sset[unicode.ToUpper(rune(s[i]))]) || (unicode.IsUpper(rune(s[i])) && sset[unicode.ToLower(rune(s[i]))])) {
			is_full = false
			off = i
			break
		}

	}
	if is_full {
		return s
	}
	res1 := LongestNiceSubstring(s[:off])
	res2 := LongestNiceSubstring(s[off+1:])
	if len(res1) >= len(res2) {
		return res1
	} else {
		return res2
	}
}
