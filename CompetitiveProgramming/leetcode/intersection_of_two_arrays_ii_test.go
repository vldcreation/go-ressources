package leetcode

import (
	"testing"

	"github.com/vldcration/go-ressources/util"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			name:  "TC1",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			expected: [][]int{
				{2, 2},
			},
		},
		{
			name:  "TC2",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			expected: [][]int{
				{4, 9},
			},
		},
		{
			name:  "TC3",
			nums1: []int{0, 2, 7, 7, 9},
			nums2: []int{0, 8, 4},
			expected: [][]int{
				{0},
			},
		},
		{
			name:  "TC4",
			nums1: []int{2, 1},
			nums2: []int{1, 1},
			expected: [][]int{
				{1},
			},
		},
		{
			name:  "TC5",
			nums1: []int{2, 1},
			nums2: []int{1, 2},
			expected: [][]int{
				{1, 2},
			},
		},
	}

	for _, tc := range tests {
		result := intersect(tc.nums1, tc.nums2)

		ok := false
		for _, v := range tc.expected {
			if util.CompareSliceInt(result, v) {
				ok = true
				break
			}
		}

		if !ok {
			t.Errorf("%s failed, Expected one of %v, got %v", tc.name, tc.expected, result)
		}
	}
}
