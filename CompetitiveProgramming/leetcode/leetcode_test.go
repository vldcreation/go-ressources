package leetcode_test

import (
	"fmt"
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode"
)

func TestMustLoadProblems(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "test 1",
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Logf("problems: %v", leetcode.NewLeetcode(leetcode.MustLoadProblems()))
		if got := len(leetcode.NewLeetcode(leetcode.MustLoadProblems()).Problems); got != test.expected {
			t.Errorf("Expected %v, but got %v", test.expected, got)
		}
	}
}

func TestZigzagConvert(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		{
			name:     "Test1",
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
	}

	for _, tt := range tests {
		fmt.Println("Len of s: ", len(tt.s))
		if got := leetcode.ZigzagConvert(tt.s, tt.numRows); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Test1",
			s:        "1337c0d3",
			expected: 1337,
		},
		{
			name:     "Test2",
			s:        "   -042",
			expected: -42,
		},
		{
			name:     "Test3",
			s:        "0-1",
			expected: 0,
		},
		{
			name:     "Test4",
			s:        "-2147483647",
			expected: -2147483647,
		},
		{
			name:     "Test5",
			s:        "20000000000000000000",
			expected: 2147483647,
		},
		{
			name:     "Test6",
			s:        "words and 987",
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := leetcode.MyAtoi__GPT(tt.s); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
