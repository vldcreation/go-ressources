package binary_tree

import (
	"testing"

	"github.com/vldcration/go-ressources/util"
)

func TestInorderTraversal(t *testing.T) {
	node1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Left: nil,
	}

	tests := []struct {
		root           *TreeNode
		expectedResult []int
	}{
		{
			root:           node1,
			expectedResult: []int{1, 3, 2},
		},
	}

	for _, test := range tests {
		result := InorderTraversal(test.root)
		if !util.CompareSliceInt(result, test.expectedResult) {
			t.Errorf("Expected %v, got %v", test.expectedResult, result)
		}
	}
}
