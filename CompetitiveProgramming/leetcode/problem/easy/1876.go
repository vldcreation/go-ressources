package easy

import "fmt"

func CountGoodSubstrings(s string) int {
	ans := 0

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			for k := j + 1; k < len(s); k++ {
				// fmt.Println(i, j, k)
				if ok(string(s[i]) + string(s[j]) + string(s[k])) {
					ans++
				}
			}
		}
	}

	return ans
}

func ok(s string) bool {
	fmt.Println(s)
	return s[0] != s[1] && s[1] != s[2] && s[0] != s[2]
}
