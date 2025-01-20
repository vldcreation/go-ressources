package podify_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/interview/podify"
)

func TestCase1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "test1",
			input:    "341234567890123",
			expected: "American Express",
		},
		{
			name:     "test2",
			input:    "341234567890123",
			expected: "American Express",
		},
		{
			name:     "test3",
			input:    "5112345678901234",
			expected: "MasterCard",
		},
		{
			name:     "test4",
			input:    "5212345678901234",
			expected: "MasterCard",
		},
		{
			name:     "test5",
			input:    "5312345678901234",
			expected: "MasterCard",
		},
		{
			name:     "test6",
			input:    "5412345678901234",
			expected: "MasterCard",
		},
		{
			name:     "test7",
			input:    "5512345678901234",
			expected: "MasterCard",
		},
		{
			name:     "test8",
			input:    "4123456789012",
			expected: "Visa",
		},
		{
			name:     "test9",
			input:    "4123456789012345",
			expected: "Visa",
		},
		{
			name:     "test10",
			input:    "4123456789012345678",
			expected: "Visa",
		},
		{
			name:     "test11",
			input:    "6011123456789012345",
			expected: "Discover",
		},
		{
			name:     "test12",
			input:    "6221261234567890",
			expected: "Discover",
		},
		{
			name:     "test13",
			input:    "6441234567890123",
			expected: "Discover",
		},
		{
			name:     "test14",
			input:    "6512345678901234",
			expected: "Discover",
		},
		{
			name:     "test15",
			input:    "5034567890123456",
			expected: "Maestro",
		},
		{
			name:     "test16",
			input:    "582126123456",
			expected: "Maestro",
		},
	}

	for _, tt := range tests {
		if got := podify.Case1(tt.input); got != tt.expected {
			t.Errorf("%s expected: %+v but got: %+v\n", tt.name, tt.expected, got)
		}
	}
}
