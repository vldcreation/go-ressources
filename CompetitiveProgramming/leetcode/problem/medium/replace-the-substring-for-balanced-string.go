package medium

func BalancedString(s string) int {
	n := len(s)
	count := make(map[byte]int)
	for i := 0; i < n; i++ {
		count[s[i]]++
	}

	if count['Q'] == n/4 && count['W'] == n/4 && count['E'] == n/4 && count['R'] == n/4 {
		return 0
	}

	ans := n
	left := 0
	for right := 0; right < n; right++ {
		count[s[right]]--
		for left <= right && count['Q'] <= n/4 && count['W'] <= n/4 &&
			count['E'] <= n/4 && count['R'] <= n/4 {
			if right-left+1 < ans {
				ans = right - left + 1
			}
			count[s[left]]++
			left++
		}
	}

	return ans
}
