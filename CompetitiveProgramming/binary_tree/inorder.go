package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  eg1
//  1
// root =
// [1,null,2,3]
// Output: [1,3,2]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	result := []int{}
	dfs(root, &result)
	return result
}

func dfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	dfs(node.Left, result)
	*result = append(*result, node.Val)
	dfs(node.Right, result)
}
