package medium

func EqualSubstring(s string, t string, maxCost int) int {
	cost := 0

	l, r := 0, 0
	for r < len(s) {
		cost += abs(int(s[r]) - int(t[r]))
		if cost > maxCost {
			cost -= abs(int(s[l]) - int(t[l]))
			l++
		}
		r++
	}
	return r - l
}
