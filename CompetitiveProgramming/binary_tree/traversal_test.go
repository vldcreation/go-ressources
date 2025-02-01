package binary_tree

import (
	"testing"

	"github.com/vldcreation/go-ressources/util"
)

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		root           *TreeNode
		expectedResult []int
	}{
		{
			root:           makeNode([]*int{toIntPtr(1), toIntPtr(2), toIntPtr(3), toIntPtr(4), toIntPtr(5), toIntPtr(6), toIntPtr(7)}),
			expectedResult: []int{4, 2, 5, 1, 6, 3, 7},
		},
	}

	for _, test := range tests {
		result := InorderTraversal(test.root)
		if !util.CompareSliceInt(result, test.expectedResult) {
			t.Errorf("Expected %v, got %v", test.expectedResult, result)
		}
	}
}

func TestPreOrderTraversal(t *testing.T) {
	tests := []struct {
		root           *TreeNode
		expectedResult []int
	}{
		{
			root:           makeNode([]*int{toIntPtr(1), toIntPtr(2), toIntPtr(3), toIntPtr(4), toIntPtr(5), toIntPtr(6), toIntPtr(7)}),
			expectedResult: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}

	for _, test := range tests {
		result := PreOrderTraversal(test.root)
		if !util.CompareSliceInt(result, test.expectedResult) {
			t.Errorf("Expected %v, got %v", test.expectedResult, result)
		}
	}

}

func TestPostOrderTraversal(t *testing.T) {
	tests := []struct {
		root           *TreeNode
		expectedResult []int
	}{
		{
			root:           makeNode([]*int{toIntPtr(1), toIntPtr(2), toIntPtr(3), toIntPtr(4), toIntPtr(5), toIntPtr(6), toIntPtr(7)}),
			expectedResult: []int{4, 5, 2, 6, 7, 3, 1},
		},
	}

	for _, test := range tests {
		result := PostOrderTraversal(test.root)
		if !util.CompareSliceInt(result, test.expectedResult) {
			t.Errorf("Expected %v, got %v", test.expectedResult, result)
		}
	}

}

func TestLevelOrderTraversal(t *testing.T) {
	tests := []struct {
		root           *TreeNode
		expectedResult []int
	}{
		{
			root:           makeNode([]*int{toIntPtr(1), toIntPtr(2), toIntPtr(3), toIntPtr(4), toIntPtr(5), toIntPtr(6), toIntPtr(7)}),
			expectedResult: []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, test := range tests {
		result := LevelOrderTraversal(test.root)
		if !util.CompareSliceInt(result, test.expectedResult) {
			t.Errorf("Expected %v, got %v", test.expectedResult, result)
		}
	}
}

func makeNode(val []*int) *TreeNode {
	if len(val) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: *val[0],
	}

	queue := []*TreeNode{root}
	i := 1
	for i < len(val) {
		node := queue[0]
		queue = queue[1:]

		if val[i] != nil {
			node.Left = &TreeNode{
				Val: *val[i],
			}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(val) && val[i] != nil {
			node.Right = &TreeNode{
				Val: *val[i],
			}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func toIntPtr(val int) *int {
	return &val
}
