package medium

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

func MyAtoi(s string) int {
	clearS := strings.TrimSpace(s)
	clearS = strings.TrimPrefix(clearS, "_")

	sign := true
	ans := ""

	if len(clearS) > 0 && clearS[0] == '-' {
		sign = false
	}

	if len(clearS) > 0 && clearS[0] == '+' {
		sign = true
	}

	for i, ch := range clearS {
		if i == 0 && (ch == '-' || ch == '+') {
			continue
		}

		if unicode.IsDigit(ch) {
			ans += string(ch)
		} else {
			break
		}
	}

	if len(ans) < 1 {
		return 0
	}

	ansConverted, err := strconv.Atoi(ans)
	if err != nil {
		if !sign {
			return -2147483648
		}

		return 2147483647
	}

	if !sign {
		ansConverted *= -1
	}

	if ansConverted <= -2147483648 {
		return -2147483648
	}

	if ansConverted >= 2147483647 {
		return 2147483647
	}

	return ansConverted
}

func MyAtoi__GPT(s string) int {
	sign := 1
	ans := 0
	i := 0

	// Skip leading whitespaces
	for i < len(s) && s[i] == ' ' {
		i++
	}

	// Check for sign
	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}

	// Convert digits to integer
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		// Check for overflow
		if ans > math.MaxInt32/10 || (ans == math.MaxInt32/10 && int(s[i]-'0') > math.MaxInt32%10) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		ans = ans*10 + int(s[i]-'0')
		i++
	}

	return ans * sign
}
