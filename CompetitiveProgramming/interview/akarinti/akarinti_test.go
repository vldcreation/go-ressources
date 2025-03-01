package akarinti_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/interview/akarinti"
)

func TestStringChallenge(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "test1",
			s:        "coderbyte",
			expected: "etybredoc",
		},
		{
			name:     "test2",
			s:        "I Love Code",
			expected: "edoC evoL I",
		},
	}

	for _, tt := range tests {
		if got := akarinti.StringChallenge(tt.s); got != tt.expected {
			t.Errorf("expected: %+v but got: %+v\n", tt.expected, got)
		}
	}
}

func TestStringChallenge2(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "test1",
			s:        "hello*3",
			expected: "Ifmmp*3",
		},
		{
			name:     "test2",
			s:        "fun times!",
			expected: "gvO Ujnft!",
		},
	}

	for _, tt := range tests {
		if got := akarinti.StringChallenge2(tt.s); got != tt.expected {
			t.Errorf("expected: %+v but got: %+v\n", tt.expected, got)
		}
	}
}

func TestStringChallenge3(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "test1",
			s:        "Hello-LOL",
			expected: "hELLO-lol",
		},
		{
			name:     "test2",
			s:        "Sup DUDE!!?",
			expected: "sUP dude!!?",
		},
	}

	for _, tt := range tests {
		if got := akarinti.StringChallenge3(tt.s); got != tt.expected {
			t.Errorf("expected: %+v but got: %+v\n", tt.expected, got)
		}
	}
}

func TestMathChallenge(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "test1",
			s:        "100101",
			expected: "37",
		},
		{
			name:     "test2",
			s:        "011",
			expected: "3",
		},
	}

	for _, tt := range tests {
		if got := akarinti.MathChallenge(tt.s); got != tt.expected {
			t.Errorf("expected: %+v but got: %+v\n", tt.expected, got)
		}
	}
}

func TestSearchinglenge(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "test1",
			s:        "Hello apple pie",
			expected: "Hello",
		},
		{
			name:     "test2",
			s:        "No words",
			expected: "-1",
		},
		{
			name:     "test3",
			s:        "I am is the best in the word haaaaaaaaaah cool bro yeaheahhheahhhhhhhhhhhhhhhh",
			expected: "yeaheahhheahhhhhhhhhhhhhhhh",
		},
	}

	for _, tt := range tests {
		if got := akarinti.SearchingChallenge(tt.s); got != tt.expected {
			t.Errorf("expected: %+v but got: %+v\n", tt.expected, got)
		}
	}
}
