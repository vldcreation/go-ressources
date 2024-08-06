package money

import "strings"

func Join(str []string, sep string) string {
	for i, s := range str {
		str[i] = strings.TrimSpace(s)
	}
	return strings.Join(str, sep)
}
