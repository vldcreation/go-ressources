package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  reference: https://www.geeksforgeeks.org/tree-traversals-inorder-preorder-and-postorder/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorder(root, &result)
	return result
}

func inorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	inorder(node.Left, result)
	*result = append(*result, node.Val)
	inorder(node.Right, result)
}

func PreOrderTraversal(root *TreeNode) []int {
	result := []int{}
	preorder(root, &result)
	return result
}

func preorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Val)
	preorder(node.Left, result)
	preorder(node.Right, result)
}

func PostOrderTraversal(root *TreeNode) []int {
	result := []int{}
	postorder(root, &result)
	return result
}

func postorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	postorder(node.Left, result)
	postorder(node.Right, result)
	*result = append(*result, node.Val)
}

func LevelOrderTraversal(root *TreeNode) []int {
	result := []int{}
	levelorder(root, &result)
	return result
}

func levelorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	queue := []*TreeNode{node}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		*result = append(*result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
