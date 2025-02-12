package bv_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/interview/bv"
)

func TestGetMaximumEfficiency(t *testing.T) {
	tests := []struct {
		name      string
		capacity  []int32
		numServer []int32
		expected  int64
	}{
		{
			name:      "test1",
			capacity:  []int32{1, 2, 3, 4},
			numServer: []int32{4},
			expected:  3,
		},
		{
			name:      "test2",
			capacity:  []int32{4, 2, 1},
			numServer: []int32{1, 1, 1},
			expected:  0,
		},
		{
			name:      "test3",
			capacity:  []int32{3, 6, 1, 2},
			numServer: []int32{1, 3},
			expected:  5,
		},
	}

	for _, tt := range tests {
		if got := bv.GetMaximumEfficiency2(tt.capacity, tt.numServer); got != tt.expected {
			t.Errorf("failed (%s) expected: %+v but got: %+v\n", tt.name, tt.expected, got)
		}
	}
}
