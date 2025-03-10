package dp_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/dsa/dp"
)

func TestNaiveFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Test1",
			n:        0,
			expected: 0,
		},
		{
			name:     "Test2",
			n:        1,
			expected: 1,
		},
		{
			name:     "Test3",
			n:        2,
			expected: 1,
		},
		{
			name:     "Test4",
			n:        3,
			expected: 2,
		},
		{
			name:     "Test5",
			n:        4,
			expected: 3,
		},
		{
			name:     "Test6",
			n:        5,
			expected: 5,
		},
		{
			name:     "Test7",
			n:        10,
			expected: 55,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := dp.NaiveFibonacci(test.n)
			if actual != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, actual)
			}
		})
	}
}

// func TestMemoizedFibonacci(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		n        int
// 		expected int
// 	}{
// 		{
// 			name:     "Test1",
// 			n:        0,
// 			expected: 0,
// 		},
// 		{
// 			name:     "Test2",
// 			n:        1,
// 			expected: 1,
// 		},
// 		{
// 			name:     "Test3",
// 			n:        2,
// 			expected: 1,
// 		},
// 		{
// 			name:     "Test4",
// 			n:        3,
// 			expected: 2,
// 		},
// 		{
// 			name:     "Test5",
// 			n:        4,
// 			expected: 3,
// 		},
// 		{
// 			name:     "Test6",
// 			n:        5,
// 			expected: 5,
// 		},
// 		{
// 			name:     "Test7",
// 			n:        10,
// 			expected: 55,
// 		},
// 	}

//		for _, test := range tests {
//			t.Run(test.name, func(t *testing.T) {
//				actual := dp.MemoizedFibonacci(test.n, make([]int, test.n+1))
//				if actual != test.expected {
//					t.Errorf("Expected %d, but got %d", test.expected, actual)
//				}
//			})
//		}
//	}
func TestTabulatedFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Test1",
			n:        0,
			expected: 0,
		},
		{
			name:     "Test2",
			n:        1,
			expected: 1,
		},
		{
			name:     "Test3",
			n:        2,
			expected: 1,
		},
		{
			name:     "Test4",
			n:        3,
			expected: 2,
		},
		{
			name:     "Test5",
			n:        4,
			expected: 3,
		},
		{
			name:     "Test6",
			n:        5,
			expected: 5,
		},
		{
			name:     "Test7",
			n:        10,
			expected: 55,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := dp.TabulatedFibonacci(test.n)
			if actual != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, actual)
			}
		})
	}
}
func TestSpaceOptimizedFibonacci(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Test1",
			n:        0,
			expected: 0,
		},
		{
			name:     "Test2",
			n:        1,
			expected: 1,
		},
		{
			name:     "Test3",
			n:        2,
			expected: 1,
		},
		{
			name:     "Test4",
			n:        3,
			expected: 2,
		},
		{
			name:     "Test5",
			n:        4,
			expected: 3,
		},
		{
			name:     "Test6",
			n:        5,
			expected: 5,
		},
		{
			name:     "Test7",
			n:        10,
			expected: 55,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := dp.SpaceOptimizedFibonacci(test.n)
			if actual != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, actual)
			}
		})
	}
}
