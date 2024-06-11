package leetcode

import (
	"testing"

	"github.com/vldcration/go-ressources/util"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected []string
	}{
		{
			name:     "Test 1",
			str:      "babad",
			expected: []string{"bab", "aba"},
		},
		{
			name:     "Test 2",
			str:      "cbbd",
			expected: []string{"bb"},
		},
		{
			name:     "Test 3",
			str:      "ccc",
			expected: []string{"ccc"},
		},
	}

	for _, tc := range tests {
		result := longestPalindrome(tc.str)

		if !util.CheckStringInSlice(result, tc.expected) {
			t.Errorf("%s failed, Expected one of %v, got %s", tc.name, tc.expected, result)
		}
	}
}
