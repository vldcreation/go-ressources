package interview

import "testing"

// col 1 = 3
// col 5 = 14
// col 18 = 18

// col = 3
// row = 5

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		col, row int
		expected int
	}{
		{
			name:     "test1",
			col:      3,
			row:      5,
			expected: 9,
		},
		{
			name:     "test2",
			col:      1,
			row:      7,
			expected: 27,
		},
		{
			name:     "test3",
			col:      5,
			row:      1,
			expected: 9,
		},
	}

	for _, tt := range tests {
		if got := Solution(tt.col, tt.row); got != tt.expected {
			t.Errorf("expected: %+v but got: %+v\n", tt.expected, got)
		}
	}
}
